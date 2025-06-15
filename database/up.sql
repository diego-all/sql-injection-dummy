--CREATE DATABASE IF NOT EXISTS sqli;

\c sqli;

DROP TABLE IF EXISTS users;
    
CREATE TABLE
  public.users (
    id serial NOT NULL,
    email character varying (255) NOT NULL,
    first_name character varying (255) NOT NULL,
    last_name character varying (255) NOT NULL,
    address character varying (255) NOT NULL,
    password character varying (60) NOT NULL,
    user_active integer NOT NULL DEFAULT 0,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
  );

  -- ALTER TABLE
  --   public.users
  -- ADD
  --   CONSTRAINT users_pkey PRIMARY KEY (id);

-- CREATE UNIQUE INDEX users_index_2 ON "public"."users" (id);


INSERT INTO public.users (email, first_name, last_name, address, password, user_active, created_at, updated_at) 
VALUES 
  ('usuario1@example.com', 'Juan', 'Pérez', 'Calle 123', '$2a$12$IjOgt/06hlNF13IOsrb8veJemUeSDB.7X27UtSubDbjBgXuL.j5ci', 1, now(), now()),
  ('usuario2@example.com', 'María', 'Gómez', 'Avenida 456', '$2a$12$xzfjUjBa06RwrNRu.wb.M.8bWJMc2cI9GZObV9495ypXRbfjNUyPS', 1, now(), now()),
  ('usuario3@example.com', 'Luis', 'Martínez', 'Calle 789', '$2a$12$1HgyDgcSZuZQDkKbEN6elug3P5Z62Rjrrf/YQdDEBiJ3sSuxcqpWW', 1, now(), now()),
  ('usuario4@example.com', 'Ana', 'Rodríguez', 'Avenida 101112', '$2a$12$IjOgt/06hlNF13IOsrb8veJemUeSDB.7X27UtSubDbjBgXuL.j5ci', 1, now(), now()),
  ('usuario5@example.com', 'Pedro', 'López', 'Calle 131415', '$2a$12$xzfjUjBa06RwrNRu.wb.M.8bWJMc2cI9GZObV9495ypXRbfjNUyPS', 1, now(), now());


