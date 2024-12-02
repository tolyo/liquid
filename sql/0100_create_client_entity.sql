-- +goose Up

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION
  create_client_entity(
    external_id_param text
  )

  -- app_entity.pub_id
  RETURNS TEXT

LANGUAGE 'plpgsql'
AS $$
DECLARE
  app_entity_instance app_entity%ROWTYPE;
BEGIN

  SELECT * FROM app_entity 
  WHERE external_id = external_id_param
  INTO app_entity_instance;

  IF FOUND THEN
    RETURN app_entity_instance.pub_id;
  END IF;

  -- create app_entity
  INSERT INTO app_entity (external_id, kind)
  VALUES (external_id_param, 'CUSTOMER')
  RETURNING * INTO app_entity_instance;

  RETURN app_entity_instance.pub_id;
END;
$$;

-- +goose StatementEnd

-- +goose Down
DROP FUNCTION IF EXISTS create_client_entity(TEXT);