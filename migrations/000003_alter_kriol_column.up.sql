-- Filename - 000003_alter_kriol_column.up.sql
ALTER TABLE entries ADD COLUMN created_at timestamp(0) with time zone NOT NULL DEFAULT NOW();
