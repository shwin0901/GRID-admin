import * as React from "react"

import { cn } from "@/lib/utils";
import { Eye, EyeOff } from "lucide-react";

function Input({ className, type, ...props }: React.ComponentProps<"input">) {
  return (
    <input
      type={type}
      data-slot="input"
      className={cn(
        "file:text-foreground placeholder:text-muted-foreground selection:bg-primary selection:text-primary-foreground dark:bg-input/30 border-input h-9 w-full min-w-0 rounded-md border bg-transparent px-3 py-1 text-base shadow-xs transition-[color,box-shadow] outline-none file:inline-flex file:h-7 file:border-0 file:bg-transparent file:text-sm file:font-medium disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
        "focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
        "aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive",
        className
      )}
      {...props}
    />
  )
}

function PasswordInput({ className, ...props }: React.ComponentProps<typeof Input>) {
  const [showPassword, setShowPassword] = React.useState(false)

  const togglePasswordVisibility = (e: React.MouseEvent) => {
    e.stopPropagation()
    setShowPassword(!showPassword)
  }

  return (
    <div className={cn("relative w-full", className)}>
      <Input
        type={showPassword ? "text" : "password"}
        autoComplete="current-password"
        className={cn("pr-10")}
        {...props}
      />
      <button
        type="button"
        onClick={togglePasswordVisibility}
        className={cn(
          "absolute right-2 top-1/2 -translate-y-1/2",
          "h-6 w-6 flex items-center justify-center",
          "text-muted-foreground hover:text-foreground transition-colors",
          "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
          "cursor-pointer bg-transparent border-0 p-0"
        )}
        aria-label={showPassword ? "hide password" : "show password"}
      >
        {showPassword ? (
          <EyeOff size={16} />
        ) : (
          <Eye size={16} />
        )}
      </button>
    </div>
  )
}

export { Input, PasswordInput }
