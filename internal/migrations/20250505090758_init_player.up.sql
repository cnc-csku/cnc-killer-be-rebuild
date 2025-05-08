CREATE TABLE players (  
    player_id UUID PRIMARY KEY,
    nickname TEXT,
    secret_code TEXT,
    victim_id UUID,
    is_alive Boolean,
    score INT,
    action_id UUID,
    email TEXT,
    
    FOREIGN KEY (action_id) REFERENCES actions(action_id),
    FOREIGN KEY (email) REFERENCES users(email)
);