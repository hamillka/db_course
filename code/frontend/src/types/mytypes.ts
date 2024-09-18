export enum ROLE {
  ADMIN,
  USER,
  DOCTOR,
}

export interface MyToken {
  exp: string;
  role: number;
}

export interface MenuItem {
  label: string;
  value: number;
}
