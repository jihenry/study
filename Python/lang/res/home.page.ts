import { Component, OnInit, OnDestroy, OnChanges, SimpleChanges, AfterViewInit, AfterContentInit } from '@angular/core';
import { Events } from '@ionic/angular';
import { AppEvent } from 'src/app/service/config/EventCode';

declare let cc: any;
declare let window: any;

@Component({
  selector: 'app-home',
  templateUrl: './home.page.html',
  styleUrls: ['./home.page.scss'],
})
export class HomePage implements OnInit, OnDestroy, OnChanges, AfterViewInit, AfterContentInit {
  cb: (ev: Event) => any;

  settingJs: HTMLScriptElement;
  mainJs: HTMLScriptElement;

  constructor(private events: Events) { }

  ngOnInit() {
    console.log("homePage ngOnInit")
    this.cb = () => {
      this.onGameSceneLanch()
    };
    document.addEventListener("GameSceneLaunch", this.cb, false);
    const myTabBar = document.getElementById('myTabBar');
    myTabBar.style.display = "none";

    this.settingJs = document.createElement('script');
    this.settingJs.src = 'src/settings.1111.js';
    document.body.appendChild(this.settingJs);

    this.mainJs = document.createElement('script');
    this.mainJs.src = 'main.de2ec.js';
    document.body.appendChild(this.mainJs);

    var cocos2d = document.createElement('script');
    cocos2d.src = 'cocos2d-js-min.e0380.js';
    var engineLoaded = () => {
      document.body.removeChild(cocos2d);
      cocos2d.removeEventListener('load', engineLoaded, false);
      window.boot();
    };
    cocos2d.addEventListener('load', engineLoaded, false);
    document.body.appendChild(cocos2d);
  }

  onGameSceneLanch() {
    cc.view.resizeWithBrowserSize(false);
    document.removeEventListener("GameSceneLaunch", this.cb, false);
    if (window.gameStatus !== 4) {
      const myTabBar = document.getElementById('myTabBar');
      myTabBar.style.display = "flex";
    }
    this.events.publish(AppEvent.app_check_addiction)
  }

  ngOnDestroy(): void {
    document.body.removeChild(this.mainJs);
    document.body.removeChild(this.settingJs);
    console.log("homepage ngOnDestroy");
  }

  ngOnChanges(changes: SimpleChanges): void {
    console.log("homepage ngOnDestroy");
  }

  ngAfterViewInit(): void {
    console.log("homepage ngAfterViewInit");
  }

  ngAfterContentInit(): void {
    console.log("homepage ngAfterContentInit");
  }
  ionViewWillEnter() {
    console.log('home 即将进入');
  }

  ionViewDidEnter() {
    console.log('home 进入之后');
  }

  ionViewWillLeave() {
    console.log('home 将要离开');
  }

  ionViewDidLeave() {
    console.log('home 离开之后');
  }

  ionViewWillUnload() {
    console.log('home 即将卸载销毁');
  }
}
