import Swal from "sweetalert2";

interface messageParams {
  message: string;
  icon?: "warning" | "error" | "question" | "success";
}

export const errorMessage = ({ message, icon = "warning" }: messageParams): void => {
  Swal.fire({
    title: message,
    timer: 1200,
    allowOutsideClick: false,
    icon,
    showConfirmButton: false,
  });
};

export const successMessage = ({ message, icon = "success" }: messageParams): void => {
  Swal.fire({
    title: message,
    timer: 1200,
    allowOutsideClick: false,
    icon,
    showConfirmButton: false,
  });
};
