CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    user_name varchar(32) UNIQUE,
    email varchar(256) UNIQUE NOT NULL,
    password varchar(100) NOT NULL,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW()
 );
 
CREATE TABLE IF NOT EXISTS task_lists (
    id uuid PRIMARY KEY,
    title varchar NOT NULL,
    description varchar,
    user_id uuid references users(id),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW()
 );

CREATE TABLE IF NOT EXISTS tasks (
    id uuid PRIMARY KEY,
    title varchar NOT NULL,
    description varchar,
    due_date date,
    task_list_id uuid references task_lists(id),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW()
 );

CREATE TABLE IF NOT EXISTS labels (
    id uuid PRIMARY KEY,
    name varchar(100) NOT NULL,
    color varchar(30),
    user_id uuid references users(id),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW()
 );

 