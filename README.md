# Team Package

The 'Team' package empowers administrators to create specialized teams with diverse users and roles tailored for distinct domains. Specific roles, assigned with precise permissions, ensure efficient collaboration and task assignment within each team. With our Teams package, administrators can effortlessly create and manage specialized teams within Golang projects, fostering efficient collaboration and task delegation across various domains.


## Features

- Team package provides functionalities to list, create, update, and delete user accounts and team members. 
- These functions ensure data integrity by validating user information such as email addresses,  phone numbers, and usernames. 
- Users can securely change their passwords, while administrators can check if specific roles are already assigned. 
- Detailed user information is retrievable, and users have the autonomy to update their own profiles. 

# Installation

``` bash
go get github.com/spurtcms/team
```


# Usage Example


```bash
func main(){
	
	//authsetup automatically migrate auth related tables in your databases=.
	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "SecretKey@123",
		RoleId: 1,
		DB: &gorm.DB{},
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD, 1)

	//Teamsetup automatically migrate team related tables in your database.
	team := teams.TeamSetup(teams.Config{
		DB:               &gorm.DB{},
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	
	if permisison {

		teamuser, count, err := team.ListUser(10, 0, teams.Filters{},1)
		//handle error
		fmt.Println(teamuser,count,err)
		
		//create user- TeamCreate struct we have multiple fields for creating user
		user,cerr:=team.CreateUser(teams.TeamCreate{FirstName: "demo", RoleId: 1, Email: "mailto:demo@gmail.com",TenantId: 1})
		//handle error
		fmt.Println(user,cerr)

		//update user- TeamCreate struct we have multiple fields for update user
		user1,uerr:=team.UpdateUser(teams.TeamCreate{FirstName: "demo1", RoleId: 2, Email: "mailto:demo1@gmail.com"}, 1,1)
		//handle error
		fmt.Println(user1,uerr)

		//delete user
		err1:=team.DeleteUser([]int{},1,1,1)
		//handle error
		fmt.Println(err1)

		
	}else{
	
		fmt.Println("unauthroized")
	}
	
}

```

# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/team/issues]. 
