INSERT INTO service_params(id, service_id, service_item_param_key_id, created_at, updated_at)
VALUES ('0a0f32dc-2629-4abc-a2a7-764fea0f1ac7', (SELECT id from re_services WHERE code = 'DPK'), (SELECT id FROM service_item_param_keys WHERE key = 'ContractCode'), now(), now());