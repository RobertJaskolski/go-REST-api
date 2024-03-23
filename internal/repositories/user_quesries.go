package repositories

var createUserQuery = `INSERT INTO users (email, first_name, last_name, time_zone, mobile, role, is_active, password) VALUES (@Email, @FirstName, @LastName, @TimeZone, @Mobile, @Role, @IsActive, @Password) RETURNING (id, email, first_name, last_name, time_zone, mobile, role, is_active, created_at, modified_at);`
var getUserQuery = `SELECT (id, email, first_name, last_name, time_zone, mobile, role, is_active, created_at, modified_at) FROM users WHERE id = @ID;`
var getLoggedQuery = `SELECT (id, email, role, is_active, password) FROM users WHERE email = @Email;`
var updateUserQuery = `UPDATE users SET email = COALESCE(@Email, email), first_name = COALESCE(@FirstName, first_name), last_name = COALESCE(@LastName, last_name), time_zone = COALESCE(@TimeZone, time_zone), mobile = COALESCE(@Mobile, mobile), role = COALESCE(@Role, role), is_active = COALESCE(@IsActive, is_active), modified_at = CURRENT_TIMESTAMP WHERE id = @ID RETURNING (id, email, first_name, last_name, time_zone, mobile, role, is_active, created_at, modified_at);`
var deleteUserQuery = `DELETE FROM users WHERE id = @ID;`
