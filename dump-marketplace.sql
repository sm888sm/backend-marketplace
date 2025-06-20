PGDMP  2                    }            marketplace    16.2 (Debian 16.2-1.pgdg120+2)    17.0 <    m           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false            n           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false            o           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false            p           1262    147866    marketplace    DATABASE     v   CREATE DATABASE marketplace WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';
    DROP DATABASE marketplace;
                     myuser    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                     pg_database_owner    false            q           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                        pg_database_owner    false    4            �            1259    148059 
   categories    TABLE     �   CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    deleted_at timestamp without time zone
);
    DROP TABLE public.categories;
       public         heap r       myuser    false    4            �            1259    148058    categories_id_seq    SEQUENCE     �   CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.categories_id_seq;
       public               myuser    false    219    4            r           0    0    categories_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;
          public               myuser    false    218            �            1259    148107    order_items    TABLE     v  CREATE TABLE public.order_items (
    id integer NOT NULL,
    order_id integer NOT NULL,
    product_id integer NOT NULL,
    quantity integer NOT NULL,
    price numeric(10,2) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);
    DROP TABLE public.order_items;
       public         heap r       myuser    false    4            �            1259    148106    order_items_id_seq    SEQUENCE     �   CREATE SEQUENCE public.order_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.order_items_id_seq;
       public               myuser    false    225    4            s           0    0    order_items_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.order_items_id_seq OWNED BY public.order_items.id;
          public               myuser    false    224            �            1259    148089    orders    TABLE     �  CREATE TABLE public.orders (
    id integer NOT NULL,
    customer_id integer NOT NULL,
    total_amount numeric(10,2) NOT NULL,
    discount_amount numeric(10,2) DEFAULT 0,
    shipping_fee numeric(10,2) DEFAULT 0,
    final_amount numeric(10,2) NOT NULL,
    status character varying(20) DEFAULT 'pending'::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone,
    CONSTRAINT orders_status_check CHECK (((status)::text = ANY ((ARRAY['pending'::character varying, 'completed'::character varying, 'cancelled'::character varying])::text[])))
);
    DROP TABLE public.orders;
       public         heap r       myuser    false    4            �            1259    148088    orders_id_seq    SEQUENCE     �   CREATE SEQUENCE public.orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.orders_id_seq;
       public               myuser    false    4    223            t           0    0    orders_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;
          public               myuser    false    222            �            1259    148126    product_images    TABLE     C  CREATE TABLE public.product_images (
    id integer NOT NULL,
    product_id integer NOT NULL,
    path character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);
 "   DROP TABLE public.product_images;
       public         heap r       myuser    false    4            �            1259    148125    product_images_id_seq    SEQUENCE     �   CREATE SEQUENCE public.product_images_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.product_images_id_seq;
       public               myuser    false    4    227            u           0    0    product_images_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.product_images_id_seq OWNED BY public.product_images.id;
          public               myuser    false    226            �            1259    148068    products    TABLE     �  CREATE TABLE public.products (
    id integer NOT NULL,
    merchant_id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    price numeric(10,2) NOT NULL,
    stock integer NOT NULL,
    category_id integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);
    DROP TABLE public.products;
       public         heap r       myuser    false    4            �            1259    148067    products_id_seq    SEQUENCE     �   CREATE SEQUENCE public.products_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.products_id_seq;
       public               myuser    false    4    221            v           0    0    products_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;
          public               myuser    false    220            �            1259    148035    schema_migrations    TABLE     c   CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);
 %   DROP TABLE public.schema_migrations;
       public         heap r       myuser    false    4            �            1259    148043    users    TABLE     -  CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    role character varying(20) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone,
    CONSTRAINT users_role_check CHECK (((role)::text = ANY ((ARRAY['merchant'::character varying, 'customer'::character varying])::text[])))
);
    DROP TABLE public.users;
       public         heap r       myuser    false    4            �            1259    148042    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public               myuser    false    4    217            w           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public               myuser    false    216            �           2604    148062    categories id    DEFAULT     n   ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);
 <   ALTER TABLE public.categories ALTER COLUMN id DROP DEFAULT;
       public               myuser    false    219    218    219            �           2604    148110    order_items id    DEFAULT     p   ALTER TABLE ONLY public.order_items ALTER COLUMN id SET DEFAULT nextval('public.order_items_id_seq'::regclass);
 =   ALTER TABLE public.order_items ALTER COLUMN id DROP DEFAULT;
       public               myuser    false    225    224    225            �           2604    148092 	   orders id    DEFAULT     f   ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);
 8   ALTER TABLE public.orders ALTER COLUMN id DROP DEFAULT;
       public               myuser    false    222    223    223            �           2604    148129    product_images id    DEFAULT     v   ALTER TABLE ONLY public.product_images ALTER COLUMN id SET DEFAULT nextval('public.product_images_id_seq'::regclass);
 @   ALTER TABLE public.product_images ALTER COLUMN id DROP DEFAULT;
       public               myuser    false    227    226    227            �           2604    148071    products id    DEFAULT     j   ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);
 :   ALTER TABLE public.products ALTER COLUMN id DROP DEFAULT;
       public               myuser    false    221    220    221            �           2604    148046    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public               myuser    false    217    216    217            b          0    148059 
   categories 
   TABLE DATA           :   COPY public.categories (id, name, deleted_at) FROM stdin;
    public               myuser    false    219   K       h          0    148107    order_items 
   TABLE DATA           t   COPY public.order_items (id, order_id, product_id, quantity, price, created_at, updated_at, deleted_at) FROM stdin;
    public               myuser    false    225   ]K       f          0    148089    orders 
   TABLE DATA           �   COPY public.orders (id, customer_id, total_amount, discount_amount, shipping_fee, final_amount, status, created_at, updated_at, deleted_at) FROM stdin;
    public               myuser    false    223   L       j          0    148126    product_images 
   TABLE DATA           b   COPY public.product_images (id, product_id, path, created_at, updated_at, deleted_at) FROM stdin;
    public               myuser    false    227   M       d          0    148068    products 
   TABLE DATA           �   COPY public.products (id, merchant_id, name, description, price, stock, category_id, created_at, updated_at, deleted_at) FROM stdin;
    public               myuser    false    221   0M       ^          0    148035    schema_migrations 
   TABLE DATA           ;   COPY public.schema_migrations (version, dirty) FROM stdin;
    public               myuser    false    215   �N       `          0    148043    users 
   TABLE DATA           h   COPY public.users (id, username, password, email, role, created_at, updated_at, deleted_at) FROM stdin;
    public               myuser    false    217   �N       x           0    0    categories_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.categories_id_seq', 3, true);
          public               myuser    false    218            y           0    0    order_items_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.order_items_id_seq', 8, true);
          public               myuser    false    224            z           0    0    orders_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.orders_id_seq', 10, true);
          public               myuser    false    222            {           0    0    product_images_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.product_images_id_seq', 1, false);
          public               myuser    false    226            |           0    0    products_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.products_id_seq', 14, true);
          public               myuser    false    220            }           0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 5, true);
          public               myuser    false    216            �           2606    148066    categories categories_name_key 
   CONSTRAINT     Y   ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_name_key UNIQUE (name);
 H   ALTER TABLE ONLY public.categories DROP CONSTRAINT categories_name_key;
       public                 myuser    false    219            �           2606    148064    categories categories_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.categories DROP CONSTRAINT categories_pkey;
       public                 myuser    false    219            �           2606    148114    order_items order_items_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.order_items DROP CONSTRAINT order_items_pkey;
       public                 myuser    false    225            �           2606    148100    orders orders_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_pkey;
       public                 myuser    false    223            �           2606    148133 "   product_images product_images_pkey 
   CONSTRAINT     `   ALTER TABLE ONLY public.product_images
    ADD CONSTRAINT product_images_pkey PRIMARY KEY (id);
 L   ALTER TABLE ONLY public.product_images DROP CONSTRAINT product_images_pkey;
       public                 myuser    false    227            �           2606    148077    products products_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.products DROP CONSTRAINT products_pkey;
       public                 myuser    false    221            �           2606    148039 (   schema_migrations schema_migrations_pkey 
   CONSTRAINT     k   ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);
 R   ALTER TABLE ONLY public.schema_migrations DROP CONSTRAINT schema_migrations_pkey;
       public                 myuser    false    215            �           2606    148057    users users_email_key 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT users_email_key;
       public                 myuser    false    217            �           2606    148053    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public                 myuser    false    217            �           2606    148055    users users_username_key 
   CONSTRAINT     W   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_username_key;
       public                 myuser    false    217            �           2606    148115 %   order_items order_items_order_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;
 O   ALTER TABLE ONLY public.order_items DROP CONSTRAINT order_items_order_id_fkey;
       public               myuser    false    225    3268    223            �           2606    148120 '   order_items order_items_product_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE;
 Q   ALTER TABLE ONLY public.order_items DROP CONSTRAINT order_items_product_id_fkey;
       public               myuser    false    225    3266    221            �           2606    148101    orders orders_customer_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES public.users(id) ON DELETE CASCADE;
 H   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_customer_id_fkey;
       public               myuser    false    223    217    3258            �           2606    148134 -   product_images product_images_product_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.product_images
    ADD CONSTRAINT product_images_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE;
 W   ALTER TABLE ONLY public.product_images DROP CONSTRAINT product_images_product_id_fkey;
       public               myuser    false    221    227    3266            �           2606    148083 "   products products_category_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE SET NULL;
 L   ALTER TABLE ONLY public.products DROP CONSTRAINT products_category_id_fkey;
       public               myuser    false    3264    221    219            �           2606    148078 "   products products_merchant_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_merchant_id_fkey FOREIGN KEY (merchant_id) REFERENCES public.users(id) ON DELETE CASCADE;
 L   ALTER TABLE ONLY public.products DROP CONSTRAINT products_merchant_id_fkey;
       public               myuser    false    3258    217    221            b   3   x�3�tK,�������2�t�I�.)�����9��sJKR�@�=... c
�      h   �   x�}��!�3V���G,"l�us�lV<>�2(800� ���Kx��=Q�Hy��&�a���_�I���,wQI
Q�?JL�9R|�ԫ�/�Y
��E%Yu%=�A��1vQI�=�<�C5J����	�kQ�e9���Y7u��Z;ZXl�����D���Z� �gn0      f   �   x���An1EמS�A��1�C�Y����z�U+3����a���I8��SO�������}|߾��v�AV�U�Q[�I��E��������V6�b.u&���n6��3�0Y�P�I��`���GF�)�	�E�L
Fݖq��������>��x`��ͧnP(�m����I�i	�{�������-�0��Y�%��3)0���D��ֽ$��,o�L6C5R?�� �J˲�2!�      j      x������ � �      d   V  x���An�0E��)|�X�x�8١.X@)+6�%$Iz�N@)$W����/��3�e�-|��E���릋�����}�w�eJ�R�T1�@����ArP�Q ؀�[E��ؗ��I��0�l��Ј.��E87���x(n�;�HOp6�30L�΅,�)I��-��s��tw�Q�:��[$�੟����I��F��·|s�fRp���z���v��4d8 ��/緛+�^0쎙��qZ�,��om�/d�^b؍K_B,�QK�!FJS��o��ᳶo5��sN����!^fJ2��D;/��a���/�i��������⁗�W��C�5I�"�NDQ��f�Q      ^      x�3�L����� �S      `   l  x����S�@�3����~�凧E�2I4i�,����H$�_3M2]���y�y�(<�����$���b>Ц��H-��a>�����Hʽ>�P��;;�F�� ��)�˯�;q�q�(���,��$W ���-�B;��P`�K�j� %�mJ]����{~���\Vο�y&�j�u�����ٓ�-�P��O6\Z�-G�u������7���N����t w �A13襨��#k��-i�j���f����{�%��؇8j�G���A8���}>�3�����������ڥ���H����o��������j4E~֖^f�ҳ
��P�����{XG���7��&���`��
�B�F�53��     