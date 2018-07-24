--
-- PostgreSQL database dump
--

-- Dumped from database version 10.4
-- Dumped by pg_dump version 10.4

-- Started on 2018-07-24 04:17:37 UTC

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 196 (class 1259 OID 16456)
-- Name: configurations; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.configurations (
    id integer NOT NULL,
    namespace text,
    version numeric DEFAULT 1,
    config_store jsonb,
    service_id numeric
);


ALTER TABLE public.configurations OWNER TO postgresdev;

--
-- TOC entry 197 (class 1259 OID 16463)
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
-- TOC entry 2186 (class 0 OID 0)
-- Dependencies: 197
-- Name: configurations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.configurations_id_seq OWNED BY public.configurations.id;


--
-- TOC entry 202 (class 1259 OID 16493)
-- Name: namespaces; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.namespaces (
    id numeric NOT NULL,
    service_id numeric NOT NULL,
    namespace text NOT NULL,
    active_version numeric
);


ALTER TABLE public.namespaces OWNER TO postgresdev;

--
-- TOC entry 198 (class 1259 OID 16465)
-- Name: services; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.services (
    id integer NOT NULL,
    owner_id numeric,
    name text NOT NULL,
    created_at timestamp(6) without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.services OWNER TO postgresdev;

--
-- TOC entry 199 (class 1259 OID 16472)
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
-- TOC entry 2187 (class 0 OID 0)
-- Dependencies: 199
-- Name: services_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.services_id_seq OWNED BY public.services.id;


--
-- TOC entry 200 (class 1259 OID 16474)
-- Name: users; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying,
    username character varying,
    password character varying,
    email character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgresdev;

--
-- TOC entry 201 (class 1259 OID 16480)
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
-- TOC entry 2188 (class 0 OID 0)
-- Dependencies: 201
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 2045 (class 2604 OID 16482)
-- Name: configurations id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.configurations ALTER COLUMN id SET DEFAULT nextval('public.configurations_id_seq'::regclass);


--
-- TOC entry 2047 (class 2604 OID 16483)
-- Name: services id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services ALTER COLUMN id SET DEFAULT nextval('public.services_id_seq'::regclass);


--
-- TOC entry 2048 (class 2604 OID 16484)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 2050 (class 2606 OID 16486)
-- Name: configurations configurations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.configurations
    ADD CONSTRAINT configurations_pkey PRIMARY KEY (id);


--
-- TOC entry 2058 (class 2606 OID 16500)
-- Name: namespaces namespaces_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.namespaces
    ADD CONSTRAINT namespaces_pkey PRIMARY KEY (id);


--
-- TOC entry 2052 (class 2606 OID 16488)
-- Name: services services_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_pkey PRIMARY KEY (id);


--
-- TOC entry 2054 (class 2606 OID 16490)
-- Name: services unique_name; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT unique_name UNIQUE (name);


--
-- TOC entry 2056 (class 2606 OID 16492)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- Completed on 2018-07-24 04:17:41 UTC

--
-- PostgreSQL database dump complete
--

