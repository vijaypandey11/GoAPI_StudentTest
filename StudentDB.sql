
CREATE TABLE Student
(
	ID INT NOT NULL,
	NAME NVARCHAR(255) NOT NULL,
	AGE INT NOT NULL,
	MARKS INT NOT NULL
)
GO

CREATE PROCEDURE GetStudentById
(
	@ID INT 
)
AS
BEGIN
	SELECT 
		ID,
		Name,
		Age,
		Marks
	FROM
		Student
	WHERE ID = @ID
END
GO

CREATE PROCEDURE GetStudentByName
(
	@Name NVARCHAR(255) 
)
AS
BEGIN
	SELECT 
		ID,
		Name,
		Age,
		Marks
	FROM
		Student
	WHERE Name = @Name
END
GO

CREATE PROCEDURE GetStudentByAge
(
	@Age INT 
)
AS
BEGIN
	SELECT 
		ID,
		Name,
		Age,
		Marks
	FROM
		Student
	WHERE Age = @Age
END
GO

CREATE PROCEDURE GetStudentByStatus
(
	@Status NVARCHAR(4) 
)
AS
BEGIN
	
	IF @Status ='pass'
	BEGIN
		SELECT 
			ID,
			Name,
			Age,
			Marks
		FROM
			Student
		WHERE Marks > 65
	END
	ELSE IF @Status ='fail'
	BEGIN
		SELECT 
			ID,
			Name,
			Age,
			Marks
		FROM
			Student
		WHERE Marks < 65
	END
END
GO

CREATE PROCEDURE GetStudentCount
AS
BEGIN
	SELECT 
		COUNT(0) AS studentCount
	FROM
		Student
	
END
GO

CREATE PROCEDURE GetAvgMarks
AS
BEGIN
	SELECT 
		AVG(MARKS) AS avgMarks
	FROM
		Student
END
GO

CREATE PROCEDURE GetMinMarks
AS
BEGIN
	SELECT 
		MIN(MARKS) AS minMarks
	FROM
		Student
END
GO

CREATE PROCEDURE GetMaxMarks
AS
BEGIN
	SELECT 
		MAX(MARKS) AS maxMarks
	FROM
		Student
END


GO