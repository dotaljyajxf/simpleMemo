import React from "react";
import { Menu } from "antd";
import { HashRouter as Router, Route, Link, Switch } from "react-router-dom";

class LeftSider extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <Router>
        <div className="left-sider">
          <Menu
            theme="light"
            mode="vertical"
            defaultSelectedKeys={["1"]}
            style={{ lineHeight: "64px" }}
          >
            <Menu.Item key="1">
              <Link to="/cal">日历</Link>
            </Menu.Item>
            <Menu.Item key="2">
              {/* 2 */}
              <Link to="/docs">文档</Link>
            </Menu.Item>
            <Menu.Item key="3">
              {/* 3 */}
              <Link to="/memo">备忘录</Link>
            </Menu.Item>
          </Menu>
        </div>
      </Router>
    );
  }
}

export default LeftSider;
