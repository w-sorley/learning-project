package wang.sorley.shiro.chapter3;

import org.apache.shiro.SecurityUtils;
import org.apache.shiro.authc.UsernamePasswordToken;
import org.apache.shiro.subject.Subject;

import static wang.sorley.shiro.chapter2.LoginLogout.testLogin;
public class RoleTest {

    public static void main(String[] args) {
        String userRoleConfgFilePath = "classpath:chapter3/shiro-role.ini";
        UsernamePasswordToken token = new UsernamePasswordToken("wang", "123");
        testLogin(token,userRoleConfgFilePath);
        Subject subject = SecurityUtils.getSubject();
        ControlBasedRole(subject);
        ControlBasedResource(subject);

        subject.logout();


    }


    /*基于角色的权限控制*/
    public static void ControlBasedRole(Subject subject){
        if(subject.hasRole("role1")){
            System.out.println("wang has role1");
        }else {
            System.out.println("wang does not have role1");
        }
        if(subject.hasRole("role2")){
            System.out.println("wang has role2");
        }else {
            System.out.println("wang does not have role2");
        }
        try {
            subject.checkRoles("role2", "role3");
            System.out.println("wang had role2 or role3");
        }catch (Exception e){
            System.out.println("wang don't had role2 or role3");
        }
    }

    public static void ControlBasedResource(Subject subject){

        if(subject.isPermitted("user:update")){
            System.out.println("role wang has permittion user:update");
        }else {
            System.out.println("role wang don't has permittion user:create");
        };
        if (subject.isPermitted("user:delete")){
            System.out.println("role wang has permittion user:delete");
        }else {
            System.out.println("role wang don't has permittion user:delete");
        }

    }
}
