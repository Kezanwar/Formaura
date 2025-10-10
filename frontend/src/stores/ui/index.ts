import { observable, action, makeObservable } from 'mobx';

type Theme = 'light' | 'dark';

class UIStore {
  theme: Theme = 'dark';
  isLoading: boolean = false;

  constructor() {
    makeObservable(this, {
      theme: observable,
      setTheme: action,
      isLoading: observable,
      setIsLoading: action
    });

    const saved = localStorage.getItem('$MobX-theme') as Theme;
    if (saved) {
      this.setTheme(saved);
    }

    document.documentElement.classList.remove('light', 'dark');
    document.documentElement.classList.add(this.theme);
  }

  setTheme(theme: Theme) {
    this.theme = theme;
    document.documentElement.classList.remove('light', 'dark');
    document.documentElement.classList.add(theme);
    localStorage.setItem('$MobX-theme', theme);
  }

  setIsLoading(isLoading: boolean) {
    this.isLoading = isLoading;
  }
}

export default UIStore;
