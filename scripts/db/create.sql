/*
USE [Master]

IF db_id('Articles') IS NULL
BEGIN
    CREATE DATABASE Articles
END
GO

USE [Articles]
*/

IF OBJECT_ID (N'Article', N'U') IS NULL 
BEGIN
    CREATE TABLE [dbo].[Article] (
        Id bigint identity(1,1) not null,
        Title nvarchar(500) not null,
        Body nvarchar(max) not null,
        CreationDate date not null,
        CONSTRAINT PK_ArticleId PRIMARY KEY CLUSTERED (Id)
    )
END
GO

IF OBJECT_ID (N'Tag', N'U') IS NULL 
BEGIN
    CREATE TABLE [dbo].[Tag] (
        Id bigint IDENTITY(1,1) not null,
        Name nvarchar(500) not null,
        CreationDate date not null,
        CONSTRAINT PK_TagId PRIMARY KEY CLUSTERED (Id)
    )
END
GO

IF OBJECT_ID (N'ArticleTag', N'U') IS NULL 
BEGIN
    CREATE TABLE [dbo].[ArticleTag] (
        ArticleId bigint not null,
        TagId bigint not null,
        CONSTRAINT PK_ArticleIdTagId PRIMARY KEY CLUSTERED (ArticleId, TagId),
        CONSTRAINT FK_ArticleId FOREIGN KEY (ArticleId) REFERENCES Article (Id),
        CONSTRAINT FK_TagId FOREIGN KEY (TagId) REFERENCES Tag (Id)
    )
END
GO

IF OBJECT_ID(N'[dbo].[spPostArticle]', N'P') IS NOT NULL
BEGIN
    DROP PROCEDURE [dbo].[spPostArticle]
END
GO

CREATE PROCEDURE [dbo].[spPostArticle]
    @Title nvarchar(500),   
    @Body nvarchar(max),
    @Tags nvarchar(max)
AS   
BEGIN
    DECLARE @ArticleId bigint
    
    INSERT INTO [dbo].[Article] VALUES (ltrim(rtrim(@Title)), ltrim(rtrim(@Body)), cast(getdate() as date))
    SET @ArticleId = SCOPE_IDENTITY()

    SELECT * INTO #TempTags FROM STRING_SPLIT(@Tags, ',')

    WHILE (SELECT Count(*) From #TempTags) > 0
    BEGIN
        DECLARE @Tag nvarchar(500)
        DECLARE @TagId bigint

        SELECT TOP 1 @Tag = ltrim(rtrim(lower(value))) From #TempTags

        IF NOT EXISTS (SELECT * FROM [dbo].[Tag] WHERE Name = @Tag)
          BEGIN
            INSERT INTO [dbo].[Tag] VALUES (@Tag, cast(getdate() as date))
            SET @TagId = SCOPE_IDENTITY()
          END
        ELSE
          BEGIN
            SELECT @TagId = Id FROM [dbo].[Tag] WHERE Name = @Tag
          END
        
        INSERT INTO [dbo].[ArticleTag] VALUES (@ArticleId, @TagId)

        Delete #TempTags Where ltrim(rtrim(lower(value))) = @Tag
    END

    DROP TABLE #TempTags
    SELECT [Id], CONVERT(VARCHAR(10), CreationDate, 112) FROM Article WHERE [Id] = @ArticleId
END
GO


IF OBJECT_ID(N'[dbo].[spGetArticles]', N'P') IS NOT NULL
BEGIN
    DROP PROCEDURE [dbo].[spGetArticles]
END
GO

CREATE PROCEDURE [dbo].[spGetArticles]
    @Id bigint null
AS   
BEGIN
    IF @Id <= 0 SET @Id = null
    SELECT 
      PT.Id, PT.Title, PT.CreationDate[Date], PT.Body, Tags = 
        STUFF (
          (
              SELECT ',' + CT.Name
              FROM [dbo].[Tag] CT
              WHERE CT.Id in (select TagId from ArticleTag where ArticleId = PT.Id)
              ORDER BY CT.Name
              FOR XML PATH(''),TYPE
          ).value('.','VARCHAR(MAX)'),1,1,SPACE(0)
        )
    FROM 
      [dbo].[Article] PT
    WHERE 
      @Id IS null OR PT.Id = @Id
    GROUP BY 
      PT.Id, PT.Title, PT.CreationDate, PT.Body
END
GO


IF OBJECT_ID(N'[dbo].[spGetTags]', N'P') IS NOT NULL
BEGIN
    DROP PROCEDURE [dbo].[spGetTags]
END
GO

CREATE PROCEDURE [dbo].[spGetTags]
    @Tag nvarchar(500),
    @Date nvarchar(50)
AS   
BEGIN
SELECT 
      PT.Id, PT.Name, 
      Articles = 
        ISNULL(STUFF (
          (
              SELECT ',' + convert(nvarchar(50), CT.Id)
              FROM [DBO].[Article] CT
              WHERE CT.Id in (select ArticleId from ArticleTag ta inner join Article a on ta.ArticleId = a.Id where ta.TagId = PT.Id AND a.CreationDate = Convert(date, @Date))
              ORDER BY CT.Id
              FOR XML PATH(''),TYPE
          ).value('.','VARCHAR(MAX)'),1,1,SPACE(0)
        ), ''),
      RelatedTags = 
        ISNULL(STUFF (
          (
              SELECT ',' + CT.Name
              FROM [DBO].[Tag] CT
              WHERE CT.Id in (select TagId from ArticleTag where ArticleId IN (select ArticleId from ArticleTag ta inner join Article a on ta.ArticleId = a.Id where ta.TagId = PT.Id AND a.CreationDate = Convert(date, @Date))) AND ltrim(rtrim(lower(CT.Name))) <> ltrim(rtrim(lower(@Tag)))
              ORDER BY CT.Id
              FOR XML PATH(''),TYPE
          ).value('.','VARCHAR(MAX)'),1,1,SPACE(0)
        ), '')
    FROM 
      [dbo].[Tag] PT
    WHERE 
      ltrim(rtrim(lower(PT.Name))) = ltrim(rtrim(lower(@Tag)))
    GROUP BY 
      PT.Id, PT.Name
END
GO