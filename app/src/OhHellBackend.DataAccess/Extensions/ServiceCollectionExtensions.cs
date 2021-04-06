using Microsoft.Extensions.DependencyInjection;
using OhHellBackend.DataAccess.Services;

namespace OhHellBackend.DataAccess.Extensions
{
    public static class ServiceCollectionExtensions
    {
        public static IServiceCollection AddDataAccess(this IServiceCollection services)
        {
            services.AddTransient<IS3Service, S3Service>();
            return services;
        }
    }
}