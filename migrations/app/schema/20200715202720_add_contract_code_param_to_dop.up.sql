INSERT INTO service_params(id, service_id, service_item_param_key_id, created_at, updated_at)
VALUES ('429a7749-f978-4054-a95d-17e8bbf64e48', (SELECT id FROM re_services WHERE code = 'DOP'), (SELECT id FROM service_item_param_keys WHERE key = 'ContractCode'), now(), now());