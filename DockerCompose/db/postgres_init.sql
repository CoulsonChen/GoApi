CREATE TABLE IF NOT EXISTS users( 
    acct varchar(20) NOT NULL, 
    pwd varchar(20) NOT NULL, 
    fullname varchar(50) NOT NULL, 
    created_at timestamptz NOT NULL DEFAULT (now()), 
    updated_at timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT PK_users PRIMARY KEY (acct)
);

INSERT INTO users(acct, pwd, fullname) VALUES
 ('CoulsonChen', 'pass.123', 'Coulson Chen');