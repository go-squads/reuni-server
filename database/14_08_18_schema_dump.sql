--
-- PostgreSQL database dump
--

-- Dumped from database version 10.4
-- Dumped by pg_dump version 10.4

-- Started on 2018-08-15 06:24:44 UTC

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
-- TOC entry 214 (class 1255 OID 24865)
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
-- TOC entry 205 (class 1259 OID 24913)
-- Name: users; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.users (
    id integer,
    name character varying,
    username character varying UNIQUE,
    password character varying,
    email character varying NOT NULL UNIQUE,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    PRIMARY KEY(id)
);


ALTER TABLE public.users OWNER TO postgresdev;

--
-- TOC entry 206 (class 1259 OID 24919)
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
-- TOC entry 200 (class 1259 OID 24886)
-- Name: organization; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.organization (
    id integer,
    name text NOT NULL UNIQUE,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    PRIMARY KEY(id)
);


ALTER TABLE public.organization OWNER TO postgresdev;

--
-- TOC entry 201 (class 1259 OID 24894)
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
-- TOC entry 202 (class 1259 OID 24896)
-- Name: organization_member; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.organization_member (
    organization_id integer NOT NULL,
    user_id integer NOT NULL,
    role text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    CONSTRAINT role_constraint CHECK ((role = ANY (ARRAY['Admin'::text, 'Developer'::text, 'Auditor'::text]))),
    PRIMARY KEY(organization_id, user_id),
    FOREIGN KEY(organization_id) REFERENCES public.organization(id) ON UPDATE CASCADE ON DELETE CASCADE
);


ALTER TABLE public.organization_member OWNER TO postgresdev;

--
-- TOC entry 203 (class 1259 OID 24903)
-- Name: services; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.services (
    organization_id integer NOT NULL,
    name text NOT NULL UNIQUE,
    created_at timestamp(6) without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    authorization_token text NOT NULL UNIQUE,
    created_by text,
    PRIMARY KEY(organization_id, name),
    FOREIGN KEY(organization_id) REFERENCES public.organization(id) ON UPDATE CASCADE ON DELETE CASCADE
);


ALTER TABLE public.services OWNER TO postgresdev;




--
-- TOC entry 198 (class 1259 OID 24877)
-- Name: namespaces; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.namespaces (
    organization_id integer NOT NULL,
    service_name text NOT NULL,
    namespace text NOT NULL,
    active_version integer DEFAULT 1,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    created_by text NOT NULL,
    PRIMARY KEY(organization_id, service_name, namespace),
    FOREIGN KEY (organization_id, service_name) REFERENCES public.services(organization_id, name) ON UPDATE CASCADE ON DELETE CASCADE
);


ALTER TABLE public.namespaces OWNER TO postgresdev;


--
-- TOC entry 196 (class 1259 OID 24866)
-- Name: configurations; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.configurations (
    organization_id integer NOT NULL,
    service_name text NOT NULL,
    namespace text NOT NULL,
    version integer DEFAULT 1,
    parent_version integer DEFAULT 0,
    config_store jsonb NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    created_by text,
    PRIMARY KEY(organization_id,service_name, namespace, version),
    FOREIGN KEY(organization_id, service_name, namespace) REFERENCES public.namespaces(organization_id, service_name, namespace) ON UPDATE CASCADE ON DELETE CASCADE

);


ALTER TABLE public.configurations OWNER TO postgresdev;


--
-- TOC entry 2953 (class 0 OID 0)
-- Dependencies: 206
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;

--
-- TOC entry 2792 (class 2604 OID 24923)
-- Name: organization id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.organization ALTER COLUMN id SET DEFAULT nextval('public.organization_id_seq'::regclass);

--
-- TOC entry 2799 (class 2604 OID 24925)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


-- Completed on 2018-08-15 06:24:48 UTC

--
-- PostgreSQL database dump complete
--

