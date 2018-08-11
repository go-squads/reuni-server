--
-- PostgreSQL database dump
--

-- Dumped from database version 10.4
-- Dumped by pg_dump version 10.4

-- Started on 2018-08-10 07:56:34 UTC

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
-- TOC entry 1 (class 3079 OID 13001)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2948 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 196 (class 1259 OID 16564)
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
-- TOC entry 197 (class 1259 OID 16573)
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
-- TOC entry 2949 (class 0 OID 0)
-- Dependencies: 197
-- Name: configurations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.configurations_id_seq OWNED BY public.configurations.id;


--
-- TOC entry 198 (class 1259 OID 16575)
-- Name: namespaces; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.namespaces (
    service_id integer NOT NULL,
    namespace text NOT NULL,
    active_version integer DEFAULT 1,
    id integer NOT NULL
);


ALTER TABLE public.namespaces OWNER TO postgresdev;

--
-- TOC entry 199 (class 1259 OID 16582)
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
-- TOC entry 2950 (class 0 OID 0)
-- Dependencies: 199
-- Name: namespaces_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.namespaces_id_seq OWNED BY public.namespaces.id;


--
-- TOC entry 200 (class 1259 OID 16584)
-- Name: organization; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.organization (
    id integer NOT NULL,
    name text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.organization OWNER TO postgresdev;

--
-- TOC entry 201 (class 1259 OID 16592)
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
-- TOC entry 2951 (class 0 OID 0)
-- Dependencies: 201
-- Name: organization_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.organization_id_seq OWNED BY public.organization.id;


--
-- TOC entry 202 (class 1259 OID 16594)
-- Name: organization_member; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.organization_member (
    organization_id integer NOT NULL,
    user_id integer NOT NULL,
    role text NOT NULL,
    created_at time without time zone DEFAULT now(),
    updated_at time without time zone DEFAULT now(),
    CONSTRAINT role_constraint CHECK ((role = ANY (ARRAY['Admin'::text, 'Developer'::text, 'Auditor'::text])))
);


ALTER TABLE public.organization_member OWNER TO postgresdev;

--
-- TOC entry 203 (class 1259 OID 16603)
-- Name: services; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.services (
    id integer NOT NULL,
    organization_id integer,
    name text NOT NULL,
    created_at timestamp(6) without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    authorization_token text NOT NULL
);


ALTER TABLE public.services OWNER TO postgresdev;

--
-- TOC entry 204 (class 1259 OID 16611)
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
-- TOC entry 2952 (class 0 OID 0)
-- Dependencies: 204
-- Name: services_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.services_id_seq OWNED BY public.services.id;


--
-- TOC entry 205 (class 1259 OID 16613)
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
-- TOC entry 206 (class 1259 OID 16619)
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
-- TOC entry 2953 (class 0 OID 0)
-- Dependencies: 206
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 2784 (class 2604 OID 16621)
-- Name: configurations id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.configurations ALTER COLUMN id SET DEFAULT nextval('public.configurations_id_seq'::regclass);


--
-- TOC entry 2786 (class 2604 OID 16622)
-- Name: namespaces id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.namespaces ALTER COLUMN id SET DEFAULT nextval('public.namespaces_id_seq'::regclass);


--
-- TOC entry 2789 (class 2604 OID 16623)
-- Name: organization id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.organization ALTER COLUMN id SET DEFAULT nextval('public.organization_id_seq'::regclass);


--
-- TOC entry 2795 (class 2604 OID 16624)
-- Name: services id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services ALTER COLUMN id SET DEFAULT nextval('public.services_id_seq'::regclass);


--
-- TOC entry 2796 (class 2604 OID 16625)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 2808 (class 2606 OID 16627)
-- Name: services auth_unique; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT auth_unique UNIQUE (authorization_token);


--
-- TOC entry 2798 (class 2606 OID 16629)
-- Name: configurations configurations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.configurations
    ADD CONSTRAINT configurations_pkey PRIMARY KEY (id);


--
-- TOC entry 2814 (class 2606 OID 16631)
-- Name: users email_unique; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT email_unique UNIQUE (email);


--
-- TOC entry 2800 (class 2606 OID 16633)
-- Name: namespaces id; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.namespaces
    ADD CONSTRAINT id PRIMARY KEY (id);


--
-- TOC entry 2802 (class 2606 OID 16635)
-- Name: organization name_unique; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.organization
    ADD CONSTRAINT name_unique UNIQUE (name);


--
-- TOC entry 2806 (class 2606 OID 16637)
-- Name: organization_member organization_member_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.organization_member
    ADD CONSTRAINT organization_member_pkey PRIMARY KEY (organization_id, user_id);


--
-- TOC entry 2804 (class 2606 OID 16639)
-- Name: organization organization_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.organization
    ADD CONSTRAINT organization_pkey PRIMARY KEY (id);


--
-- TOC entry 2810 (class 2606 OID 16641)
-- Name: services services_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_pkey PRIMARY KEY (id);


--
-- TOC entry 2812 (class 2606 OID 16643)
-- Name: services unique_name; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT unique_name UNIQUE (name);


--
-- TOC entry 2816 (class 2606 OID 16645)
-- Name: users username_unique; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT username_unique UNIQUE (username);


--
-- TOC entry 2818 (class 2606 OID 16647)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2819 (class 2606 OID 24756)
-- Name: services organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON UPDATE CASCADE ON DELETE CASCADE;


-- Completed on 2018-08-10 07:56:38 UTC

--
-- PostgreSQL database dump complete
--

