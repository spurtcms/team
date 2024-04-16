# Teams Package

The 'Teams' package empowers administrators to create specialized teams with diverse users and roles tailored for distinct domains. Specific roles, assigned with precise permissions, ensure efficient collaboration and task assignment within each team. With our Teams package, administrators can effortlessly create and manage specialized teams within Golang projects, fostering efficient collaboration and task delegation across various domains.


## Features

- Teams package provides functionalities to list, create, update, and delete user accounts and team members. 
- These functions ensure data integrity by validating user information such as email addresses,  phone numbers, and usernames. 
- Users can securely change their passwords, while administrators can check if specific roles are already assigned. 
- Detailed user information is retrievable, and users have the autonomy to update their own profiles. 

# Installation

``` bash
go get github.com/spurtcms/teams
```


# Usage Example


```bash
import(
	"fmt"
	"github.com/spurtcms/team"
	"github.com/spurtcms/auth"
	teamrole "github.com/spurtcms/team-roles"
)

func main(){
	
	//authsetup automatically migrate auth related tables in your databases=.
	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  SecretKey,
	})

	token, _ := Auth.CreateToken()

	//RoleSetup automatically migrate role&permission related tables in your database.
	Permission := teamrole.RoleSetup(role.Config{
		RoleId: 1,
		DB:     &gorm.DB{},
	})
	
	//Teamsetup automatically migrate team related tables in your database.
	team := TeamSetup(team.Config{
		DB:               &gorm.DB{},
		AuthEnable:       true,
		PermissionEnable: true,
		Authenticate:     auth.Authentication{Token: token, SecretKey: SecretKey},
		PermissionConf:   Permission,
		Auth:             Auth,
	})
	
	//check if login user or given roleid have this module permission?
	flg, _ := Permission.IsGranted("teams", role.CRUD)

	if flg{

		teamuser, count, err := team.ListUser(10, 0, Filters{})
		//handle error
		fmt.Println(teamuser,count,err)
		
		//create user- TeamCreate struct we have multiple fields for creating user
		user,cerr:=team.CreateUser(team.TeamCreate{FirstName: "demo", RoleId: 1, Email: "demo@gmail.com"})
		//handle error
		fmt.Println(user,cerr)

		//update user- TeamCreate struct we have multiple fields for update user
		user1,uerr:=team.UpdateUser(team.TeamCreate{FirstName: "demo1", RoleId: 2, Email: "demo1@gmail.com"}, 1)
		//handle error
		fmt.Println(user1,uerr)

		//delete user
		err1:=team.DeleteUser(1)
		//handle error
		fmt.Println(err1)

		
	}else{
	
		fmt.Println("unauthroized")
	}
	
}

```

# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/team/issues]. 
