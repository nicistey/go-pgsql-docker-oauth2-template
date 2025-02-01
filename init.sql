-- создание таблицы пользователей
CREATE TABLE IF NOT EXISTS users (
    IDus SERIAL PRIMARY KEY,  
    IDGoogle VARCHAR(255) UNIQUE, 
    name VARCHAR(255) NOT NULL,  
    email VARCHAR(255) UNIQUE NOT NULL  
);

-- создание таблицы мероприятий
CREATE TABLE IF NOT EXISTS events (
    IDev SERIAL PRIMARY KEY,  
    IDus INTEGER NOT NULL REFERENCES users(IDus),
    Event_name VARCHAR(255) NOT NULL,
    Event_time TIMESTAMPTZ, 
    Description TEXT,
    Location TEXT, 
    Is_public BOOLEAN NOT NULL 
);

-- вставка одного тестового пользователя
INSERT INTO users (name, email)
VALUES ( 'Ivan Ivanov', 'ivanov@example.com');

-- вставка одного тестового публичного мероприятия для этого пользователя
INSERT INTO events (IDus, Event_name, Event_time, Description, Location, Is_public)
VALUES (1, 'Тестовое мероприятие №1', '2023-10-10 12:00:00+03', 'Описание мероприятия', 'Место проведения', true);