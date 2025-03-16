ALTER TABLE files 
ADD COLUMN user_profile_id UUID UNIQUE NULL,
ADD CONSTRAINT fk_user_profile FOREIGN KEY (user_profile_id) REFERENCES users(id) ON DELETE SET NULL;
