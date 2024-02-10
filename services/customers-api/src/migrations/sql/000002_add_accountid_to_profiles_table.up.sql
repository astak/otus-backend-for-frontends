ALTER TABLE profiles
ADD COLUMN account_id VARCHAR(37) NOT NULL;
ALTER TABLE profiles
ADD CONSTRAINT account_id_unique UNIQUE(account_id);