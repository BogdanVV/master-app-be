DO $$ BEGIN
    CREATE TYPE day_of_week AS ENUM ('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE priority AS ENUM ('high', 'medium', 'low');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    active_days day_of_week[],
    priority priority DEFAULT 'medium',
    is_daily boolean DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL
);

-- insert into todos (
    -- title,
    -- description,
    -- active_days,
    -- priority,
    -- is_daily,
    -- created_at,
    -- updated_at,
    -- user_id
    -- ) VALUES (
    
    -- 'todo 1',
    -- 'todo 1 description',
    -- ARRAY['Monday']::day_of_week[],
    -- 'medium',
    -- false,
    -- CURRENT_TIMESTAMP,
    -- CURRENT_TIMESTAMP,
    -- 'd9b64dbd-38e9-45d3-8c25-0ba199a9bcb5'
    -- );

-- SELECT u.id,
-- u.name,
-- u.email,
-- u.created_at,
-- u.updated_at,
-- t.id,
-- t.title,
-- t.description,
-- t.active_days,
-- t.priority,
-- t.is_daily,
-- t.created_at,
-- t.updated_at
-- FROM users u
-- JOIN todos t ON u.id = t.user_id
-- WHERE u.id = 'd9b64dbd-38e9-45d3-8c25-0ba199a9bcb5';