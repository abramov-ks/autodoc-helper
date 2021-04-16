

--
-- Name: partnumber_checklist; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.partnumber_checklist (
                                             id integer NOT NULL,
                                             partnumber character varying(255) NOT NULL,
                                             inital_price numeric(8,2),
                                             date_last_checked timestamp without time zone,
                                             actual boolean DEFAULT true NOT NULL,
                                             name character varying(255)
);


ALTER TABLE public.partnumber_checklist OWNER TO postgres;

--
-- Name: partnumber_checklist_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.partnumber_checklist_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.partnumber_checklist_id_seq OWNER TO postgres;

--
-- Name: partnumber_checklist_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.partnumber_checklist_id_seq OWNED BY public.partnumber_checklist.id;


--
-- Name: partnumber_prices; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.partnumber_prices (
                                          id integer NOT NULL,
                                          partnumber character varying(255),
                                          date_checked timestamp with time zone,
                                          minimal_price numeric(8,2),
                                          info jsonb
);


ALTER TABLE public.partnumber_prices OWNER TO postgres;

--
-- Name: partnumber_prices_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.partnumber_prices_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.partnumber_prices_id_seq OWNER TO postgres;

--
-- Name: partnumber_prices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.partnumber_prices_id_seq OWNED BY public.partnumber_prices.id;


--
-- Name: partnumber_checklist id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.partnumber_checklist ALTER COLUMN id SET DEFAULT nextval('public.partnumber_checklist_id_seq'::regclass);


--
-- Name: partnumber_prices id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.partnumber_prices ALTER COLUMN id SET DEFAULT nextval('public.partnumber_prices_id_seq'::regclass);


--
-- Name: partnumber_checklist partnumber_checklist_partnumber_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.partnumber_checklist
    ADD CONSTRAINT partnumber_checklist_partnumber_key UNIQUE (partnumber);


--
-- Name: partnumber_checklist partnumber_checklist_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.partnumber_checklist
    ADD CONSTRAINT partnumber_checklist_pkey PRIMARY KEY (id);


--
-- Name: partnumber_prices partnumber_prices_id_idx; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.partnumber_prices
    ADD CONSTRAINT partnumber_prices_id_idx PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--
