/*
    DROP TABLE [dbo].[ArticleTag]
    DROP TABLE [dbo].[Article]
    DROP TABLE [dbo].[Tag]
*/
SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
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

IF TYPE_ID (N'[dbo].[TY_Tags]') IS NULL 
BEGIN
    CREATE TYPE [dbo].[TY_Tags] AS TABLE(
        Name varchar(500)  NULL
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
    @Tags [dbo].[TY_Tags] READONLY
AS   
BEGIN
    DECLARE @ArticleId bigint
    
    INSERT INTO [dbo].[Article] VALUES (ltrim(rtrim(@Title)), ltrim(rtrim(@Body)), cast(getdate() as date))
    SET @ArticleId = SCOPE_IDENTITY()

    SELECT * INTO #TempTags FROM @Tags

    WHILE (SELECT Count(*) From #TempTags) > 0
    BEGIN
        DECLARE @Tag nvarchar(500)
        DECLARE @TagId bigint

        SELECT TOP 1 @Tag = ltrim(rtrim(lower(Name))) From #TempTags

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

        Delete #TempTags Where Name = @Tag
    END

    DROP TABLE #TempTags
END
GO

/*

DECLARE @Tags AS [dbo].[TY_Tags]
INSERT INTO @Tags VALUES ('abc'), ('pqr'), ('xyz')

EXEC [dbo].[spPostArticle] 'x', 'y', @Tags

SELECT * FROM [dbo].[Article]
SELECT * FROM [dbo].[Tag]
SELECT * FROM [dbo].[ArticleTag]

*/