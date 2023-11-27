CREATE TABLE IF NOT EXISTS videos (
  id TEXT PRIMARY KEY NOT NULL,
  video_type TEXT NOT NULL,
  video_title TEXT NOT NULL,
  channel_id TEXT NOT NULL,
  FOREIGN KEY(channel_id) REFERENCES channels(id)
);
CREATE TABLE IF NOT EXISTS channels (
  id TEXT PRIMARY KEY NOT NULL,
  channel_name TEXT NOT NULL,
  channel_url TEXT NOT NULL UNIQUE,
  is_subbed BOOLEAN NOT NULL
);
CREATE TABLE IF NOT EXISTS watch_history (
  id INTEGER PRIMARY KEY,
  video_id TEXT,
  watched_at TEXT NOT NULL,
  channel_id TEXT,
  FOREIGN KEY (channel_id) REFERENCES channels(id)
  FOREIGN KEY(video_id) REFERENCES videos(id)
);
CREATE INDEX _25be8e9c140a414d88cae1490ca9cc77 ON videos (video_type);
CREATE INDEX _c4d70e462d114758899d737ef0304a59 ON channels (is_subbed);
CREATE INDEX _617adedd5c104cf6bd5020be36aed4e4 ON watch_history (video_id);