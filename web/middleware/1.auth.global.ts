export default defineNuxtRouteMiddleware((to, _) => {
  if (to.fullPath === "/auth/signin") return;

  const authStore = useAuthStore();
  if (!authStore.isLoggedIn) return navigateTo("/auth/signin");

  const projectsStore = useProjectsStore();
  projectsStore.fetch();
});
