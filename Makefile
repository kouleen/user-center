thrift=thrift
idl_path = idl/user-center

gen_service:
	kitex -module github.com/kouleen/user-center -service userheader.service $(idl_path)/user.thrift
