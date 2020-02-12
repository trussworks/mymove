-- This migration allows a faux devlocal certicate to have access to the Orders API locally.
-- Navy devlocal certificate config/tls/devlocal-faux-navy-orders.cer
INSERT INTO public.client_certs (
	id,
	sha256_digest,
	subject,
	allow_orders_api,
	allow_air_force_orders_read,
	allow_air_force_orders_write,
	allow_army_orders_read,
	allow_army_orders_write,
	allow_coast_guard_orders_read,
	allow_coast_guard_orders_write,
	allow_marine_corps_orders_read,
	allow_marine_corps_orders_write,
	allow_navy_orders_read,
	allow_navy_orders_write,
	created_at,
	updated_at)
VALUES (
	'20b085a3-12f3-4d25-9313-50b378d51305',
	'9eae654ac64507553de0d964262afc28527a461978ffc960e777f6f35942bd8f',
	'CN=localhost,OU=Not Navy Orders,O=Not Navy,L=Washington,ST=DC,C=US',
	true,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	true,
	true,
	now(),
	now());
