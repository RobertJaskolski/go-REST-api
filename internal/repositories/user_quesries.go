package repositories

var createUserQuery = `INSERT INTO users (email, first_name, last_name, time_zone, mobile, role, is_active, password) VALUES (@Email, @FirstName, @LastName, @TimeZone, @Mobile, @Role, @IsActive, @Password) RETURNING (id, email, first_name, last_name, time_zone, mobile, role, is_active, created_at, modified_at);`
var getUserQuery = `SELECT (id, email, first_name, last_name, time_zone, mobile, role, is_active, created_at, modified_at) FROM users WHERE id = @ID;`
