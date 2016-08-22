from modules import utils
from modules.logger import log
from unittest import TextTestRunner

import sys
import unittest


class PsutilCollectorLargeTest(unittest.TestCase):

    def setUp(self):
        # set and download required binaries (snapd, snapctl, plugins)
        self.binaries = utils.set_binaries()
        utils.download_binaries(self.binaries)

        log.debug("Starting snapd")
        self.binaries.snapd.start()
        if not self.binaries.snapd.isAlive():
            self.fail("snapd thread died")

        log.debug("Waiting for snapd to finish starting")
        if not self.binaries.snapd.wait():
            log.error("snapd errors: {}".format(self.binaries.snapd.errors))
            self.binaries.snapd.kill()
            self.fail("snapd not ready, timeout!")

    def test_psutil_collector_plugin(self):
        # load psutil collector
        loaded = self.binaries.snapctl.load_plugin("snap-plugin-collector-meminfo")
        self.assertTrue(loaded, "meminfo collector loaded")

        # check available metrics, plugins and tasks
        metrics = self.binaries.snapctl.list_metrics()
        plugins = self.binaries.snapctl.list_plugins()
        tasks = self.binaries.snapctl.list_tasks()
        self.assertGreater(len(metrics), 0, "Metrics available {} expected {}".format(len(metrics), 0))
        self.assertEqual(len(plugins), 1, "Plugins available {} expected {}".format(len(plugins), 1))
        self.assertEqual(len(tasks), 0, "Tasks available {} expected {}".format(len(tasks), 0))

        # check config policy for metric
        rules = self.binaries.snapctl.metric_get("/intel/procfs/meminfo/mem_total")
        self.assertEqual(len(rules), 0, "Rules available {} expected {}".format(len(rules), 0))

        # create and list available task
        task_id = self.binaries.snapctl.create_task("/snap-plugin-collector-meminfo/scripts/docker/large/meminfo-task.yml")
        tasks = self.binaries.snapctl.list_tasks()
        self.assertEqual(len(tasks), 1, "Tasks available {} expected {}".format(len(tasks), 1))

        # check if task hits and fails
        hits = self.binaries.snapctl.task_hits_count(task_id)
        fails = self.binaries.snapctl.task_fails_count(task_id)
        self.assertGreater(hits, 0, "Task hits {} expected {}".format(hits, ">0"))
        self.assertEqual(fails, 0, "Task fails {} expected {}".format(fails, 0))

        # stop task and list available tasks
        stopped = self.binaries.snapctl.stop_task(task_id)
        self.assertTrue(stopped, "Task stopped")
        tasks = self.binaries.snapctl.list_tasks()
        self.assertEqual(len(tasks), 1, "Tasks available {} expected {}".format(len(tasks), 1))

        # unload plugin, list metrics and plugins
        self.binaries.snapctl.unload_plugin("collector", "meminfo", "3")
        metrics = self.binaries.snapctl.list_metrics()
        plugins = self.binaries.snapctl.list_plugins()
        self.assertEqual(len(metrics), 0, "Metrics available {} expected {}".format(len(metrics), 0))
        self.assertEqual(len(plugins), 0, "Plugins available {} expected {}".format(len(plugins), 0))

        # check for snapd errors
        self.assertEqual(len(self.binaries.snapd.errors), 0, "Errors found during snapd execution")

    def tearDown(self):
        log.debug("Stopping snapd thread")
        self.binaries.snapd.stop()
        if self.binaries.snapd.isAlive():
            log.warn("snapd thread did not died")

if __name__ == "__main__":
    test_suite = unittest.TestLoader().loadTestsFromTestCase(PsutilCollectorLargeTest)
    test_result = TextTestRunner().run(test_suite)
    # exit with return code equal to number of failures
    sys.exit(len(test_result.failures))


