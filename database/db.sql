DROP TABLE IF EXISTS campaign_images;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS campaigns;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	occupation VARCHAR(100) NOT NULL,
	email VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	avatar_file_name VARCHAR(255),
	role VARCHAR(10),
	created_at DATE,
	updated_at DATE
);

CREATE TABLE IF NOT EXISTS campaigns(
	id SERIAL PRIMARY KEY,
	user_id INTEGER,
	name VARCHAR(255) NOT NULL,
	short_description VARCHAR(255),
	long_description TEXT,
	goal_amount INTEGER,
	current_amount INTEGER,
	backer_count INTEGER,
	perks TEXT,
	slug VARCHAR(100),
	created_at DATE,
	updated_at DATE,
	FOREIGN KEY(user_id) 
	  REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS campaign_images(
	id SERIAL PRIMARY KEY,
	campaign_id INTEGER,
	file_name VARCHAR(255),
	is_primary BOOLEAN,
	created_at DATE,
	updated_at DATE,
	FOREIGN KEY(campaign_id) 
	  REFERENCES campaigns(id)
);

CREATE TABLE IF NOT EXISTS transactions(
	id SERIAL PRIMARY KEY,
	user_id INTEGER,
	campaign_id INTEGER,
	amount INTEGER,
	status VARCHAR(255),
	code VARCHAR(255),
	created_at DATE,
	updated_at DATE,
	FOREIGN KEY(user_id) 
	  REFERENCES users(id),
	FOREIGN KEY(campaign_id) 
	  REFERENCES campaigns(id)
);

-- INSERT INTO users VALUES 
-- 	(1, 'Faris', 'Programmer', 'punyanyaarii@gmail.com', '1234', 'avatar.jpg', 'admin', NOW(), NOW()),
-- 	(2, 'Firdaus', 'Programmer', 'farisfirdausapr@gmail.com', '1234', 'avatar-2.jpg', 'user', NOW(), NOW());