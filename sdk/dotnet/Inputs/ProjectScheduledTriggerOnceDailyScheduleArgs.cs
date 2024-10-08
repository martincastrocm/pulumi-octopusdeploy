// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class ProjectScheduledTriggerOnceDailyScheduleArgs : global::Pulumi.ResourceArgs
    {
        [Input("daysOfWeeks", required: true)]
        private InputList<string>? _daysOfWeeks;

        /// <summary>
        /// The days of the week to run the trigger.
        /// </summary>
        public InputList<string> DaysOfWeeks
        {
            get => _daysOfWeeks ?? (_daysOfWeeks = new InputList<string>());
            set => _daysOfWeeks = value;
        }

        /// <summary>
        /// The time of day to start the trigger.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public ProjectScheduledTriggerOnceDailyScheduleArgs()
        {
        }
        public static new ProjectScheduledTriggerOnceDailyScheduleArgs Empty => new ProjectScheduledTriggerOnceDailyScheduleArgs();
    }
}
