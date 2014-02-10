'use strict';

var Homepage = function() {
  this.navbarName = element(by.className('name'));

  this.get = function() {
    browser.get('/');
  };
};

var Register = function() {
  this.get = function() {
    browser.get('/register');
  };
};

describe('homepage', function() {
  it('should have the hackedu name', function() {
    var homepage = new Homepage();
    homepage.get();

    expect(homepage.navbarName.getText()).toEqual('hackEDU');
  });
});

describe('register', function() {
  it('should contain login')
});
