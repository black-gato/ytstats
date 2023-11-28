CREATE TABLE videos (
  id TEXT PRIMARY KEY NOT NULL,
  video_type TEXT NOT NULL,
  video_title TEXT NOT NULL,
  channel_id TEXT,
  FOREIGN KEY(channel_id) REFERENCES channels(id)
);
CREATE TABLE channels (
  id TEXT PRIMARY KEY NOT NULL,
  channel_name TEXT NOT NULL,
  channel_url TEXT NOT NULL UNIQUE,
  is_subbed INTEGER NOT NULL
);
CREATE TABLE watch_history (
  id INTEGER PRIMARY KEY,
  video_id TEXT,
  watched_at TEXT NOT NULL,
  channel_id TEXT,
  FOREIGN KEY(channel_id) REFERENCES channels(id),
  FOREIGN KEY(video_id) REFERENCES videos(id)
);
CREATE INDEX _e2db485ef680414f9216cd60f84ae0c5 ON videos (video_type);
CREATE INDEX _cd5c39677ae24c20b90cc1510bbb0ef6 ON channels (is_subbed);
CREATE INDEX _77074f8c135647038a7ab94d6bfa14bf ON watch_history (video_id);
CREATE INDEX _dae0b1c768dd470c9f475aef92f93f14 ON watch_history (channel_id);