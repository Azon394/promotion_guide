from django.urls import path
from . import views

urlpatterns = [
    path('', views.main),
    path('registration', views.reg),
    path('auth', views.auth),
    path('account', views.account)
]