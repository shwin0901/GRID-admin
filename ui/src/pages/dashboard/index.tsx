import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { useProfile, useLogout } from '@/hooks/useAuth'

export default function DashboardPage() {
  const { data: user, isLoading } = useProfile()
  const logout = useLogout()

  if (isLoading) {
    return (
      <div className="flex min-h-svh items-center justify-center">
        <p className="text-muted-foreground">Loading...</p>
      </div>
    )
  }

  return (
    <div className="bg-muted flex min-h-svh flex-col items-center justify-center p-6">
      <Card className="w-full max-w-md">
        <CardContent className="p-6">
          <div className="flex flex-col gap-4">
            <div className="text-center">
              <h1 className="text-2xl font-bold">Dashboard</h1>
              <p className="text-muted-foreground mt-2">
                Welcome, {user?.username || 'User'}!
              </p>
            </div>
            {user && (
              <div className="rounded-lg bg-muted p-4 text-sm">
                <p><strong>ID:</strong> {user.id}</p>
                <p><strong>Username:</strong> {user.username}</p>
                <p><strong>Email:</strong> {user.email || 'Not set'}</p>
              </div>
            )}
            <Button variant="outline" onClick={logout}>
              Logout
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}
