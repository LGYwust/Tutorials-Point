import { Movie } from "@/type/movietype";
import axiosClient from "./axios";

export const Apilogin = (loginForm: {
  RoleID: number;
  Account: string;
  Password: string;
  captcha: string;
}) => {
  return axiosClient.post("/User/login", loginForm);
};

export const Apiregister = (registerForm: {
  Account: string;
  Email: string;
  Phone: string;
  Password: string;
  Confirmpassword: string;
  Captcha: string;
}) => {
  return axiosClient.post("/User/register", registerForm);
};

export const Apisearchmovies = (data: string) => {
  return axiosClient.post("/User/getsearchmovie", data);
};

export const ApiGetUserinfo = () => {
  return axiosClient.get("/User/getinfo");
};

export const ApisaveUserinfo = (data: {
  nickname: string;
  time: string;
  sex: string;
  desc: string;
}) => {
  return axiosClient.post("/User/userinfo", data);
};

export const ApichangeAvatar = (formData: FormData) => {
  return axiosClient.post("/image/upload", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
};

export const ApichangeMovie = (formData: FormData) => {
  return axiosClient.post("/image/uploadmovie", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
};

export const ApichangeConstation = (formData: FormData) => {
  return axiosClient.post("/image/uploadconsultation", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
};

export const Apiselectmovie = (active: string[]) => {
  return axiosClient.post("image/getselectmovie", active);
};

export const ApigetAvatar = () => {
  return axiosClient.get("/image/getimage");
};

export const ApigetHotMovies = () => {
  return axiosClient.get("/image/gethotmovies");
};

export const ApigetRecommendMovies = () => {
  return axiosClient.get("/image/getrecommendmovies");
};

export const ApigetAllMovies = () => {
  return axiosClient.get("/image/getallmovies");
};

export const ApigetConsultation = () => {
  return axiosClient.get("/image/getconsultation");
};

export const ApieditUser = (userForm: {
  Account: string;
  Email: string;
  Phone: string;
  RoleID: string;
}) => {
  return axiosClient.post("/User/edituser", userForm);
};

export const ApigetAllUser = () => {
  return axiosClient.get("/User/getalluser");
};

export const ApideltUser = (account: string) => {
  return axiosClient.post("/User/deltuser", account);
};

export const ApiaddMovie = (movie: {
  name: string;
  image: string;
  director: string;
  time: string;
  country: string;
  category: string;
  quote: string;
}) => {
  return axiosClient.post("/movie/addmovie", movie);
};

export const ApideltMovie = (name : string)=>{
  return axiosClient.post("/movie/deltmovie", name);
}

export const ApiaddConsultation = (consultation: {
  image: string,
  title: string,
  description: string,
  source: string,
  releaseDate: string,
  url: string,
}) => {
  return axiosClient.post("/movie/addconsultation", consultation);
};

export const ApideltConsultation = (title : string)=>{
  return axiosClient.post("/movie/deltconsultation", title);
}