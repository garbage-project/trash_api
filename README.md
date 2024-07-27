# admin-api
admin-api



```
create table trash.stock
(
    t_id                 bigint                  not null comment 'PK, Auto_incremtn',
    name                 varchar(100)            not null comment '시설물 명',
    location             varchar(255)            not null comment '주소',
    management           varchar(255) default '' null comment '관리 시설',
    detail_location      varchar(100) default '' null comment '상세 위치',
    description          varchar(300) default '' null comment '추가 설명',
    updated_by           varchar(100) default '' null comment '업데이트 주체',
    roles                int                     not null comment '구분',
    hardness             double                  not null comment '경도',
    latitude             double                  not null comment '위도',
    created_date         bigint                  not null comment '생성 날짜',
    updated_date         bigint       default 0  null comment '업데이트 날짜',
    related_phone_number varchar(20)  default '' null comment '관리 부서 전화번호',
    is_valid             tinyint(1)   default 0  not null comment '노출 여부',
    constraint unique_name_location
        unique (name, location)
);
```