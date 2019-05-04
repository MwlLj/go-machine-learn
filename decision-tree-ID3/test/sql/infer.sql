#namespace infer_db

#create tables
/*
create table if not exists t_data (
    data_uuid varchar(64),
    lable_kv json,
    classify varchar(128),
    create_time datetime,
    update_time datetime
);

create table if not exists t_tree (
    tree json
);
*/
#end

/*
    @bref 新增数据
    @is_brace true
    @in_isarr true
    @out_isarr false
    @in dataUuid: string
    @in lableKV: string
    @in classify: string
    @in createTime: string
    @in updateTime: string
*/
#define addData
insert into t_data values({0}, {1}, {2}, {3}, {4});
#end

/*
    @bref 更新数据
    @is_brace true
    @in_isarr true
    @out_isarr false
    @in dataUuid: string
    @in lableKV: string
    @in classify: string
    @in updateTime: string
*/
#define updateData
update t_data set lable_kv = {1}, classify = {2}, update_time = {3}
where data_uuid = {0};
#end

/*
    @bref 更新然后添加数据
    @is_brace true
    @in_isarr true
    @out_isarr false
    @in dataUuid: string
    @in lableKV: string
    @in classify: string
    @in updateTime: string
    @sub addData[1]
*/
#define updateThenAddData
update t_data set lable_kv = {1}, classify = {2}, update_time = {3}
where data_uuid = {0};
#end

/*
    @bref 获取全部的数据
    @is_brace true
    @in_isarr false
    @out_isarr true
    @out dataUuid: string
    @out lableKV: string
    @out classify: string
*/
#define getAllData
select data_uuid, lable_kv, classify from t_data;
#end

/*
    @bref 获取一条数据
    @is_brace true
    @in_isarr false
    @out_isarr false
    @out lableKV: string
    @out classify: string
*/
#define getOneData
select lable_kv, classify from t_data
limit 1;
#end

/*
    @bref 更新树
    @is_brace true
    @in_isarr false
    @out_isarr false
    @in tree: string
*/
#define updateTree
delete from t_tree;
insert into t_tree values({0});
#end

/*
    @bref 获取树
    @is_brace true
    @in_isarr false
    @out_isarr false
    @out tree: string
*/
#define getTree
select * from t_tree;
#end

