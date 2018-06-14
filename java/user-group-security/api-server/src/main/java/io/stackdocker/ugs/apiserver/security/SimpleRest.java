package io.stackdocker.ugs.apiserver.security;

import org.apache.shiro.SecurityUtils;
import org.apache.shiro.subject.Subject;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;


@RestController
public class SimpleRest {

    @RequestMapping("/role/config")
    public ResponseEntity roleConfig() {
        return ResponseEntity.ok("Here you are");
    }

    @RequestMapping("/user/addition")
    public ResponseEntity userAddition() {
        Subject subject = SecurityUtils.getSubject();
        boolean permitted = subject.isPermitted("system:admin:create");
        System.out.println("permitted = " + permitted);
        return ResponseEntity.ok("Here you are");
    }
}
