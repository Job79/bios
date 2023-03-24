INSERT INTO classification VALUES (1, '23012dfa-c067-443c-99ab-9397007a8cff', 'G', 'General audiences');
INSERT INTO classification VALUES (2, '4fabbcfa-f181-49cb-b730-589414c02447', 'PG', 'Parental guidance suggested');
INSERT INTO classification VALUES (3, '1636442e-9661-4d12-bba5-49afc4176ca7', 'PG-13', 'Parents strongly cautioned.');
INSERT INTO classification VALUES (4, 'e0c20441-6869-437a-9f06-377a5c1ee835', 'R', 'Restricted for children younger than 17 years without parents');
INSERT INTO classification VALUES (5, '9b19aaca-6e74-42be-b02b-702eb86c0d6e', 'NC-17', 'No one under 17 years permitted');

INSERT INTO file VALUES (1, 'dc75a270-9c1f-47f4-8c9c-c490b3a36458', '/files/dc75a270-9c1f-47f4-8c9c-c490b3a36458.webp', 'image');
INSERT INTO file VALUES (2, '33be6b70-85b5-4bc5-b15c-d10502316036', '/files/33be6b70-85b5-4bc5-b15c-d10502316036.webp', 'image');
INSERT INTO file VALUES (3, 'febe2ec1-81c3-4e95-921c-b7f43f8fc771', '/files/febe2ec1-81c3-4e95-921c-b7f43f8fc771.webp', 'image');

INSERT INTO genre VALUES (1, 'ae03261d-8cbf-4424-b0b4-fcf2477ef142', 'fantasy');
INSERT INTO genre VALUES (2, '3a3d9913-df17-477a-b479-eaa0df7be9fb', 'drama');
INSERT INTO genre VALUES (3, '688b1b8c-f813-426e-905e-188430637b36', 'thriller');
INSERT INTO genre VALUES (4, 'ffb7345e-9dd8-443d-a81f-27feeb9e995f', 'horror');
INSERT INTO genre VALUES (5, 'f9ed54a3-cf12-4235-9977-5aaaabc473fa', 'comedy');

INSERT INTO movie VALUES (1, '2c839106-d7b9-4cb6-a6e9-06671c15f407', 'Doctor strange in the multiverse of madness', 'In Marvel Studios'' Doctor Strange in the Multiverse of Madness onthult het MCU de Multiverse en verlegt het zijn grenzen verder dan ooit tevoren.', 127, '2022-05-04 19:40:55', 'now');
INSERT INTO movie VALUES (2, '89c7a94e-0c25-4a2a-8749-4e1e0e49cd83', 'costa!!', 'Spaanse stranden, zon, zee, feesten, liefde en vriendschap… met al deze ingrediënten is Costa!! dé ultieme zomerse feelgoodfilm.', 120, '2022-04-28 19:41:47', 'now');
INSERT INTO movie VALUES (3, '001f1564-c6eb-487a-904b-0949b6a23c4d', 'foodies', 'Foodies is een verfrissende en sprankelende romantische komedie. Wat heb je over voor je droombaan en de zoektocht naar ware liefde? Heb je de moed om jezelf te durven zijn?', 91, '2022-05-12 19:42:50', 'soon');


INSERT INTO movie_classification VALUES (1, 2);
INSERT INTO movie_classification VALUES (2, 3);
INSERT INTO movie_classification VALUES (3, 1);

INSERT INTO movie_file VALUES (2, 2, 0);
INSERT INTO movie_file VALUES (1, 1, 0);
INSERT INTO movie_file VALUES (3, 3, 0);

INSERT INTO movie_genre VALUES (1, 1);
INSERT INTO movie_genre VALUES (2, 2);
INSERT INTO movie_genre VALUES (3, 3);

INSERT INTO room VALUES (2, '137e1b7a-294c-4c3d-8347-4bccfef88a98', '2');
INSERT INTO room VALUES (1, 'e8a25559-b172-4595-89ac-06ef09993339', '1');

INSERT INTO showing VALUES (1, '0c881de1-ccdb-427e-8810-f81ecfe4acb7', 1, 1, '2022-05-12 19:46:03', '2022-05-12 19:46:06');
INSERT INTO showing VALUES (2, 'b9b39eba-b154-4a0f-aa60-d4a7a170b118', 2, 2, '2022-05-12 19:46:14', '2022-05-12 19:46:17');

-- password = '123456'
INSERT INTO "user" VALUES (1, '3c05b9d0-c22b-478e-b4e0-5033f952b8fa', 'admin', '\x00100800000000800000000806bcf01be42e3bc9373275f719c20a8863b64a6a89a7a727297c46ebfa1e66d4e22ae97b24ae02a2aa64753cdd3076f5');

SELECT pg_catalog.setval('classification_id_seq', 1, false);
SELECT pg_catalog.setval('file_id_seq', 3, true);
SELECT pg_catalog.setval('genre_id_seq', 1, false);
SELECT pg_catalog.setval('movie_classification_movie_id_seq', 1, false);
SELECT pg_catalog.setval('movie_id_seq', 3, true);
SELECT pg_catalog.setval('room_id_seq', 1, false);
SELECT pg_catalog.setval('showing_id_seq', 2, true);
SELECT pg_catalog.setval('user_id_seq', 1, false);
