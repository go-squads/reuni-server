--
-- PostgreSQL database dump
--

-- Dumped from database version 10.4
-- Dumped by pg_dump version 10.4

-- Started on 2018-07-24 06:24:50 UTC

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
-- TOC entry 202 (class 1259 OID 16493)
-- Name: namespaces; Type: TABLE; Schema: public; Owner: postgresdev
--

CREATE TABLE public.namespaces (
    service_id numeric NOT NULL,
    namespace text NOT NULL,
    active_version numeric DEFAULT 1,
    id integer NOT NULL
);


ALTER TABLE public.namespaces OWNER TO postgresdev;

--
-- TOC entry 203 (class 1259 OID 16505)
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
-- TOC entry 2172 (class 0 OID 0)
-- Dependencies: 203
-- Name: namespaces_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgresdev
--

ALTER SEQUENCE public.namespaces_id_seq OWNED BY public.namespaces.id;


--
-- TOC entry 2043 (class 2604 OID 16507)
-- Name: namespaces id; Type: DEFAULT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.namespaces ALTER COLUMN id SET DEFAULT nextval('public.namespaces_id_seq'::regclass);


--
-- TOC entry 2045 (class 2606 OID 16515)
-- Name: namespaces id; Type: CONSTRAINT; Schema: public; Owner: postgresdev
--

ALTER TABLE ONLY public.namespaces
    ADD CONSTRAINT id PRIMARY KEY (id);


-- Completed on 2018-07-24 06:24:52 UTC

--
-- PostgreSQL database dump complete
--

