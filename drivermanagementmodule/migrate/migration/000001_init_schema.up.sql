  CREATE TABLE IF NOT EXISTS "event" (
  "id" serial PRIMARY KEY,
  "title" varchar,
  "description" varchar
);

INSERT INTO "event" ("id","title", "description") VALUES
  (1,'Event 1', 'Description of Event 1'),
  (2,'Event 2', 'Description of Event 2'),
  (3,'Event 3', 'Description of Event 3');

CREATE TABLE IF NOT EXISTS "notes" (
  "id" integer generated always as identity,
  "subject" varchar,
  "description" varchar
);

INSERT INTO "notes" ("subject", "description") VALUES
  ('Meeting with Client', 'Discuss project requirements and timelines with the client.'),
  ('Research and Development', 'Conduct research on the latest technologies for project development.'),
  ('Team Collaboration', 'Collaborate with team members to enhance project collaboration and communication.'),
  ('Task Prioritization', 'Prioritize tasks for the upcoming sprint based on project goals.'),
  ('Bug Fixing Session', 'Identify and fix bugs reported during the testing phase.'),
  ('Training Workshop', 'Conduct a workshop to train team members on new tools and technologies.'),
  ('Monthly Review Meeting', 'Review project progress and discuss any challenges faced during the month.'),
  ('Documentation Update', 'Update project documentation to reflect recent changes and improvements.');
