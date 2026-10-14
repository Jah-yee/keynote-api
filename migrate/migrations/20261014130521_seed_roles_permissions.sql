INSERT INTO public.entity_kinds
(id, "name", "key", created_at, updated_at, deleted_at)
VALUES('295c558f-0674-46df-aed4-5ccde9fba239'::uuid, 'Users', 'users', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL);
INSERT INTO public.entity_kinds
(id, "name", "key", created_at, updated_at, deleted_at)
VALUES('9cd834d7-0574-4201-9f5a-6eb3f23ff4bf'::uuid, 'Tenant', 'tenant', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL);
INSERT INTO public.entity_kinds
(id, "name", "key", created_at, updated_at, deleted_at)
VALUES('a81f1391-426c-440e-8f7b-cd77ce88b931'::uuid, 'Workspace', 'workspace', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL);
INSERT INTO public.entity_kinds
(id, "name", "key", created_at, updated_at, deleted_at)
VALUES('0ecab892-0d6d-446d-9ab5-7d4fea9a4dd3'::uuid, 'Project', 'project', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL);


INSERT INTO public.state_definitions
(id, "key", "name", created_at, updated_at, deleted_at, entity_id)
VALUES('664b2db7-b6fa-4bdc-8be7-4b47272d5f9e'::uuid, 'active', 'Active', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL, '295c558f-0674-46df-aed4-5ccde9fba239'::uuid);
INSERT INTO public.state_definitions
(id, "key", "name", created_at, updated_at, deleted_at, entity_id)
VALUES('44ea4a85-e250-462c-bfb4-dc14efd75713'::uuid, 'deleted', 'Deleted', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL, '295c558f-0674-46df-aed4-5ccde9fba239'::uuid);
INSERT INTO public.state_definitions
(id, "key", "name", created_at, updated_at, deleted_at, entity_id)
VALUES('537a902b-6b8b-4c11-8ec8-805ffcaed75c'::uuid, 'archived', 'Archived', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL, '295c558f-0674-46df-aed4-5ccde9fba239'::uuid);
INSERT INTO public.state_definitions
(id, "key", "name", created_at, updated_at, deleted_at, entity_id)
VALUES('7f7e76af-6359-4199-a2f3-1a3302f80112'::uuid, 'active', 'Active', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL, '9cd834d7-0574-4201-9f5a-6eb3f23ff4bf'::uuid);
INSERT INTO public.state_definitions
(id, "key", "name", created_at, updated_at, deleted_at, entity_id)
VALUES('4a60dc6f-5f6d-44d7-bb6a-e5e9cbd68ce7'::uuid, 'deleted', 'Deleted', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL, '9cd834d7-0574-4201-9f5a-6eb3f23ff4bf'::uuid);
INSERT INTO public.state_definitions
(id, "key", "name", created_at, updated_at, deleted_at, entity_id)
VALUES('f4686af6-0523-4007-a295-bada3c3aaafd'::uuid, 'archived', 'Archived', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL, '9cd834d7-0574-4201-9f5a-6eb3f23ff4bf'::uuid);
INSERT INTO public.state_definitions
(id, "key", "name", created_at, updated_at, deleted_at, entity_id)
VALUES('7ccd00b8-e708-45f6-b3f5-955ecc597694'::uuid, 'active', 'Active', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL, 'a81f1391-426c-440e-8f7b-cd77ce88b931'::uuid);
INSERT INTO public.state_definitions
(id, "key", "name", created_at, updated_at, deleted_at, entity_id)
VALUES('8af31ae3-453a-4c1c-af74-7c486a6e8702'::uuid, 'deleted', 'Deleted', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL, 'a81f1391-426c-440e-8f7b-cd77ce88b931'::uuid);
INSERT INTO public.state_definitions
(id, "key", "name", created_at, updated_at, deleted_at, entity_id)
VALUES('1ba4db36-6270-4c1c-b807-02394403bfbc'::uuid, 'archived', 'Archived', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL, 'a81f1391-426c-440e-8f7b-cd77ce88b931'::uuid);
INSERT INTO public.state_definitions
(id, "key", "name", created_at, updated_at, deleted_at, entity_id)
VALUES('b7d042e8-4739-46a3-bb6f-61b9000188b9'::uuid, 'active', 'Active', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL, '0ecab892-0d6d-446d-9ab5-7d4fea9a4dd3'::uuid);
INSERT INTO public.state_definitions
(id, "key", "name", created_at, updated_at, deleted_at, entity_id)
VALUES('c339575c-2c6f-4565-ba44-fdbfda91a1bb'::uuid, 'deleted', 'Deleted', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL, '0ecab892-0d6d-446d-9ab5-7d4fea9a4dd3'::uuid);
INSERT INTO public.state_definitions
(id, "key", "name", created_at, updated_at, deleted_at, entity_id)
VALUES('7cd9c2f4-0453-4d98-af58-9a7653236ecf'::uuid, 'archived', 'Archived', '2026-10-14 18:26:50.728', '2026-10-14 18:26:50.728', NULL, '0ecab892-0d6d-446d-9ab5-7d4fea9a4dd3'::uuid);