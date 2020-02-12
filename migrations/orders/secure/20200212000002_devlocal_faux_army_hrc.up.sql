-- This migration allows a faux devlocal certicate to have access to the Orders API locally.
-- Army HRC devlocal certificate config/tls/devlocal-faux-army-hrc-orders.cer
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
	'd45f2ae5-ca66-4157-b342-82a39bb588f6',
	'ad6f04a6f6938909d201ebfe9efd1edce6e5c280f6d2f93d93bb5666cc869277',
	'CN=localhost,OU=Not Army HRC Orders,O=Not Army HRC,L=Washington,ST=DC,C=US',
	true,
	false,
	false,
	true,
	true,
	false,
	false,
	false,
	false,
	false,
	false,
	now(),
	now());
