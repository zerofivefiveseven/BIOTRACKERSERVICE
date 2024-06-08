
-- name: AddUserCredentials :one
INSERT INTO users_credentials (email, password)
VALUES (@user_email::TEXT, @user_password::TEXT)
RETURNING id;

-- name: CheckUserExists :one
SELECT EXISTS(SELECT email FROM users_credentials WHERE email = @email::TEXT);

-- name: AddUserInfo :exec
INSERT INTO users_info (id, first_name, last_name, dateOfBirthday)
VALUES (@id::INT, @first_name::TEXT, @last_name::TEXT, @dateOfBirhday::TIMESTAMP);

-- name: GetUserCredentials :one
SELECT password from users_credentials
where email = @email::TEXT;

-- name: AddPlot :exec
INSERT INTO plots (user_id, name, content)
VALUES (@user_id::INT, @name::TEXT, @content);

-- name: GetUserPlotsInfo :many
SELECT id, name from plots
where user_id = @user_id::INT;

-- name: GetPlotsByIds :many
SELECT * from plots
where id = any (@plot_ids::INT[]);
