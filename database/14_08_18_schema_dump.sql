--
-- PostgreSQL database dump
--

-- Dumped from database version 10.4
-- Dumped by pg_dump version 10.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- Name: update_timestamp(); Type: FUNCTION; Schema: public; Owner: postgresdev
--

CREATE FUNCTION public.update_timestamp() RETURNS timestamp without time zone
    LANGUAGE sql
    AS $$
select now()::timestamp
$$;


ALTER FUNCTION public.update_timestamp() OWNER TO postgresdev;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: configurations; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.configurations (
    id integer NOT NULL,
    namespace text,
    version integer DEFAULT 1,
    config_store jsonb,
    service_id integer,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    created_by text
);


ALTER TABLE public.configurations OWNER TO postgresdev;

--
-- Name: configurations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgresdev
--

CREATE SEQUENCE public.configurations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.configurations_id_seq OWNER TO postgresdev;

--
-- Name: configurations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.configurations_id_seq OWNED BY public.configurations.id;


--
-- Name: namespaces; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.namespaces (
    service_id integer NOT NULL,
    namespace text NOT NULL,
    active_version integer DEFAULT 1,
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    created_by text NOT NULL
);


ALTER TABLE public.namespaces OWNER TO postgresdev;

--
-- Name: namespaces_id_seq; Type: SEQUENCE; Schema: public; Owner: postgresdev
--

CREATE SEQUENCE public.namespaces_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.namespaces_id_seq OWNER TO postgresdev;

--
-- Name: namespaces_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.namespaces_id_seq OWNED BY public.namespaces.id;


--
-- Name: organization; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.organization (
    id integer NOT NULL,
    name text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.organization OWNER TO postgresdev;

--
-- Name: organization_id_seq; Type: SEQUENCE; Schema: public; Owner: postgresdev
--

CREATE SEQUENCE public.organization_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.organization_id_seq OWNER TO postgresdev;

--
-- Name: organization_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.organization_id_seq OWNED BY public.organization.id;


--
-- Name: organization_member; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.organization_member (
    organization_id integer NOT NULL,
    user_id integer NOT NULL,
    role text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    CONSTRAINT role_constraint CHECK ((role = ANY (ARRAY['Admin'::text, 'Developer'::text, 'Auditor'::text])))
);


ALTER TABLE public.organization_member OWNER TO postgresdev;

--
-- Name: services; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.services (
    id integer NOT NULL,
    organization_id integer,
    name text NOT NULL,
    created_at timestamp(6) without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    authorization_token text NOT NULL,
    created_by text
);


ALTER TABLE public.services OWNER TO postgresdev;

--
-- Name: services_id_seq; Type: SEQUENCE; Schema: public; Owner: postgresdev
--

CREATE SEQUENCE public.services_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.services_id_seq OWNER TO postgresdev;

--
-- Name: services_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.services_id_seq OWNED BY public.services.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying,
    username character varying,
    password character varying,
    email character varying NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgresdev;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgresdev
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgresdev;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: configurations id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.configurations ALTER COLUMN id SET DEFAULT nextval('public.configurations_id_seq'::regclass);


--
-- Name: namespaces id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.namespaces ALTER COLUMN id SET DEFAULT nextval('public.namespaces_id_seq'::regclass);


--
-- Name: organization id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.organization ALTER COLUMN id SET DEFAULT nextval('public.organization_id_seq'::regclass);


--
-- Name: services id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services ALTER COLUMN id SET DEFAULT nextval('public.services_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: configurations; Type: TABLE DATA; Schema: public; Owner: postgresdev
--

COPY public.configurations (id, namespace, version, config_store, service_id, created_at, updated_at, created_by) FROM stdin;
7	default	1	{"port": "8000"}	9	2018-08-14 07:54:13.987943	2018-08-14 07:54:13.987943	\N
8	default	1	{"port": "8080"}	8	2018-08-14 07:54:45.04702	2018-08-14 07:54:45.04702	\N
9	notdef	1	{"host": "locahost"}	8	2018-08-14 07:56:05.108715	2018-08-14 07:56:05.108715	\N
10	default	2	{"port": "8000"}	9	2018-08-14 09:08:09.167496	2018-08-14 09:08:09.167496	\N
11	default	3	{"host": "localhost", "port": "8000"}	9	2018-08-14 09:09:32.59237	2018-08-14 09:09:32.59237	\N
12	default	4	{"host": "localhost", "port": "8000", "balck": "black"}	9	2018-08-14 09:11:09.329722	2018-08-14 09:11:09.329722	\N
13	default	5	{"host": "localhost", "balck": "black"}	9	2018-08-14 09:12:42.700238	2018-08-14 09:12:42.700238	\N
\.


--
-- Data for Name: namespaces; Type: TABLE DATA; Schema: public; Owner: postgresdev
--

COPY public.namespaces (service_id, namespace, active_version, id, created_at, updated_at, created_by) FROM stdin;
8	default	1	8	2018-08-14 07:54:45.04405	2018-08-14 07:54:45.04405	rifkiadrn
8	notdef	1	9	2018-08-14 07:56:05.106501	2018-08-14 07:56:05.106501	rifkiadrn
9	default	5	7	2018-08-14 07:54:13.982079	2018-08-14 07:54:13.982079	rifkiadrn
\.


--
-- Data for Name: organization; Type: TABLE DATA; Schema: public; Owner: postgresdev
--

COPY public.organization (id, name, created_at, updated_at) FROM stdin;
11	testing-1-org	2018-08-14 07:50:13.000644	2018-08-14 07:50:13.000644
\.


--
-- Data for Name: organization_member; Type: TABLE DATA; Schema: public; Owner: postgresdev
--

COPY public.organization_member (organization_id, user_id, role, created_at, updated_at) FROM stdin;
11	52	Admin	2018-08-14 07:50:13.01519	2018-08-14 07:50:13.01519
\.


--
-- Data for Name: services; Type: TABLE DATA; Schema: public; Owner: postgresdev
--

COPY public.services (id, organization_id, name, created_at, updated_at, authorization_token, created_by) FROM stdin;
8	11	service-1-testingorg	2018-08-14 07:50:36.273075	2018-08-14 07:50:36.273075	MvpZHS1EDC7YUzACCmAvlsBXHngkNZNAOGPEDT7FtgmgIZ+exFLqJgN6A3iboI3v	rifkiadrn
9	11	service-2-testingorg	2018-08-14 07:54:13.922857	2018-08-14 07:54:13.922857	XYJCrcEo44Eb4rwMkstlD9ebHi9o4pJRqbc6L1fvvw+T5aPsEgOI15/HyfdkcHIi	rifkiadrn
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgresdev
--

COPY public.users (id, name, username, password, email, created_at, updated_at) FROM stdin;
48	kenneth	kenneth	R2B43NEMlMtG6q5f8f_5HOpM4if0eTu1CdrUb4pZaAQ=	kenneth@gmail.com	2018-08-13 09:53:47.401668	2018-08-13 09:53:47.401668
52	Rifki Adrian	rifkiadrn	4hG2ykzuMNpCx1MwO6uAz0-qT51cWKRWxYbg75YvoDU=	rifkiadrn@gmail.com	2018-08-13 12:19:16.867421	2018-08-13 12:19:16.867421
\.


--
-- Name: configurations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgresdev
--

SELECT pg_catalog.setval('public.configurations_id_seq', 13, true);


--
-- Name: namespaces_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgresdev
--

SELECT pg_catalog.setval('public.namespaces_id_seq', 9, true);


--
-- Name: organization_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgresdev
--

SELECT pg_catalog.setval('public.organization_id_seq', 11, true);


--
-- Name: services_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgresdev
--

SELECT pg_catalog.setval('public.services_id_seq', 9, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgresdev
--

SELECT pg_catalog.setval('public.users_id_seq', 52, true);


--
-- Name: services auth_unique; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT auth_unique UNIQUE (authorization_token);


--
-- Name: configurations configurations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.configurations
    ADD CONSTRAINT configurations_pkey PRIMARY KEY (id);


--
-- Name: users email_unique; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT email_unique UNIQUE (email);


--
-- Name: namespaces id; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.namespaces
    ADD CONSTRAINT id PRIMARY KEY (id);


--
-- Name: organization name_unique; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.organization
    ADD CONSTRAINT name_unique UNIQUE (name);


--
-- Name: organization_member organization_member_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.organization_member
    ADD CONSTRAINT organization_member_pkey PRIMARY KEY (organization_id, user_id);


--
-- Name: organization organization_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.organization
    ADD CONSTRAINT organization_pkey PRIMARY KEY (id);


--
-- Name: services services_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_pkey PRIMARY KEY (id);


--
-- Name: services unique_name; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT unique_name UNIQUE (name);


--
-- Name: users username_unique; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT username_unique UNIQUE (username);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: services organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--