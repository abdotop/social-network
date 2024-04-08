CREATE TABLE notifications_temp AS
SELECT id, user_id, type, message, created_at, deleted_at
FROM notifications;

-- Drop the original table
DROP TABLE notifications;

-- Rename the temporary table to the original table name
ALTER TABLE notifications_temp RENAME TO notifications;
