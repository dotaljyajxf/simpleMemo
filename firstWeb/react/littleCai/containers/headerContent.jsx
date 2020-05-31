import React, { Component } from "react";
import { Button, Icon, message } from "antd";

const zh = {
  madeBy: "嘿哈小组",
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
          <Icon type="heart" theme="filled" />
          <span>{zh.madeBy}</span>
        </div>
        <Button type="primary" onClick={this.loginHandle}>{zh.login}</Button>
      </div>
    );
  }
}

export default HeaderContent;
