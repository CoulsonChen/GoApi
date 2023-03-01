CREATE TABLE IF NOT EXISTS users( 
    acct varchar(20) NOT NULL, 
    pwd varchar(20) NOT NULL, 
    fullname varchar(50) NOT NULL, 
    created_at timestamptz NOT NULL DEFAULT (now()), 
    updated_at timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT PK_users PRIMARY KEY (acct)
);

INSERT INTO users(acct, pwd, fullname) VALUES
 ('CoulsonChen', 'pass.123', 'Coulson Chen'),
 ('CoulsonChen1', 'pass_1', 'Coulson Chen - 1'),
 ('CoulsonChen2', 'pass_2', 'Coulson Chen - 2'),
 ('CoulsonChen3', 'pass_3', 'Coulson Chen - 3'),
 ('CoulsonChen4', 'pass_4', 'Coulson Chen - 4'),
 ('CoulsonChen5', 'pass_5', 'Coulson Chen - 5'),
 ('CoulsonChen6', 'pass_6', 'Coulson Chen - 6'),
 ('CoulsonChen7', 'pass_7', 'Coulson Chen - 7');