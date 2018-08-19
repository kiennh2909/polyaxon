# Generated by Django 2.0.8 on 2018-08-19 18:05

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('db', '0005_node_scheduling'),
    ]

    operations = [
        migrations.AlterField(
            model_name='experiment',
            name='build_job',
            field=models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, related_name='experiments', to='db.BuildJob'),
        ),
        migrations.AlterField(
            model_name='job',
            name='build_job',
            field=models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, related_name='jobs', to='db.BuildJob'),
        ),
    ]
