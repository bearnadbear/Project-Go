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
	description TEXT,
	perks TEXT,
	backer_count INTEGER,
	goal_amount INTEGER,
	current_amount INTEGER,
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
	is_primary SMALLINT,
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

INSERT INTO users VALUES 
	(1, 'Faris', 'Programmer', 'punyanyaarii@gmail.com', '12345678', '', 'admin', NOW(), NOW()),
	(2, 'Firdaus', 'Programmer', 'farisfirdausapr@gmail.com', '12345678', '', 'admin', NOW(), NOW());

INSERT INTO campaigns VALUES 
	(1, 1, 'Faris-c', 'Hahihuheho', 'Huhahihuhe', 'Haha, Hihi, Huhu', 100, 10, 100, 'Hoho', NOW(), NOW()),
	(2, 2, 'Firdaus-c', 'Hahihuheho', 'Huhahihuhe', 'Hehe, Hoho, Haho', 200, 20, 200, 'Hoho', NOW(), NOW());

INSERT INTO campaign_images VALUES
	(1, 1, 'campaign-images/satu.jpg', 0, NOW(), NOW()),
	(2, 1, 'campaign-images/dua.jpg', 0, NOW(), NOW()),
	(3, 1, 'campaign-images/tiga.jpg', 1, NOW(), NOW()),
	(4, 2, 'campaign-images/satu.jpg', 0, NOW(), NOW()),
	(5, 2, 'campaign-images/dua.jpg', 1, NOW(), NOW()),
	(6, 2, 'campaign-images/tiga.jpg', 2, NOW(), NOW());