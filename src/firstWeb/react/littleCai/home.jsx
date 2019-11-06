import React, { Component } from "react";
import { Layout } from 'antd';
import 'antd/dist/antd.css';
import './style/common.css';
import './style/home.css';

import HeaderContent from './containers/headerContent.jsx';
import Body from './containers/body.jsx';
import FooterContent from './containers/footerContent.jsx';

const { Header, Footer, Content } = Layout;
const zh = {
    hello: '菜菜',
    login: '登录',
}

class HomePage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
    };
  }

  render() {
    return (
        <div className="little-cai">
          <Layout>
            <Header>
              <HeaderContent />
            </Header>
            <Content>
              <Body />
            </Content>
            <Footer>
              <FooterContent />
            </Footer>
          </Layout>
        </div>
    );
  }
}

HomePage.propTypes = {
};

export default HomePage;
