import React, { Component } from "react";
import { Button, Avatar, message } from "antd";

const zh = {
  hello: "Hello",
  login: "登录",
  message: '请输入信息'
};

class HeaderContent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
    this.loginHandle = this.loginHandle.bind(this);
  }

  loginHandle() {
    return message.info(zh.message);
  }

  render() {
    return (
      <div className="header-content">
        <div className="user-info">
          <span>{zh.hello}</span>
          {/* <Avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" /> */}
        </div>
        <Button type="primary" onClick={this.loginHandle}>{zh.login}</Button>
      </div>
    );
  }
}

export default HeaderContent;
